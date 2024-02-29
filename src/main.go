package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

type Server interface {
	IsAlive() bool
	Address() string
    Server(rw http.ResponseWriter, req *http.Request)
}

type simpleServer struct {
	addr  string
	proxy *httputil.ReverseProxy
}

func NewLoadBalancer(port string, servers []Server) *LoadBalancer {
	return &LoadBalancer{
		roundRobinCount: 0,
		port:            port,
		servers:          servers,
	}
}

func newSimpleServer(addr string) *simpleServer {
	serverUrl, err := url.Parse(addr)
	handleErr(err)

	return &simpleServer{
		addr:  addr,
		proxy: httputil.NewSingleHostReverseProxy(serverUrl),
	}
}

type LoadBalancer struct {
	port              string
	roundRobinCount   int
	servers            []Server
	currentlyServing int 
}

func handleErr(err error) {
	if err != nil {
		fmt.Printf("error: %+v\n", err)
		os.Exit(1)
	}
}

func (s *simpleServer) Address() string { return s.addr}
func (s *simpleServer) IsAlive() bool {return true}
func (s *simpleServer) Server(rw http.ResponseWriter, req *http.Request){
    s.proxy.ServeHTTP(rw, req)}


// getNextAvailableServer implements a round-robin strategy that
// considers previously marked unavailable servers.
func (lb *LoadBalancer) getNextAvailableServer() Server{
	// Start from the server after the currently serving one.
	startIndex := (lb.currentlyServing + 1) % len(lb.servers)

	for i := 0; i < len(lb.servers); i++ {
		serverIndex := (startIndex + i) % len(lb.servers)
		server := lb.servers[serverIndex]

		if server.IsAlive() {
			lb.currentlyServing = serverIndex
			return server
		}
	}

	// If no healthy server found, return nil.
	return nil
}

func (lb *LoadBalancer) serveProxy(rw http.ResponseWriter, req *http.Request) {
	server := lb.getNextAvailableServer()
	if server == nil {
		fmt.Println("No healthy server found")
		return
	}

	fmt.Printf("Forwarding request to address %q\n", server.Address())
	server.Server(rw, req)
}

func main() {
	servers := []Server{
		newSimpleServer("https://www.bing.com/"),
		newSimpleServer("https://www.facebook.com/"),
		newSimpleServer("https://duckduckgo.com/"),
		newSimpleServer("https://google.com/"),
	}

	lb := NewLoadBalancer("8000", servers)
	handleRedirect := func(rw http.ResponseWriter, req *http.Request) {
		lb.serveProxy(rw, req)
	}
	http.HandleFunc("/", handleRedirect)
	fmt.Printf("Serving requests at localhost:%s\n", lb.port)
	http.ListenAndServe(":"+lb.port, nil)
}
