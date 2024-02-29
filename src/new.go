 package main

// import (
// 	"fmt"
// 	"net/http"
// 	"net/http/httputil"
// 	"net/url"
// 	"os"
// )

// type Server interface {
// 	IsAlive() bool
// 	Address() string
// 	Server(rw http.ResponseWriter, r *http.Request)
// }

// type simpleServer struct {
// 	addr  string
// 	proxy *httputil.ReverseProxy
// }

// func NewLoadBalancer(port string, servers []Server) *LoadBalancer {
// 	return &LoadBalancer{
// 		roundRobinCount: 0,
// 		port:            port,
// 		servers:         servers,
// 	}

// }

// func newSimpleServer(addr string) *simpleServer {
// 	serverUrl, err := url.Parse(addr)
// 	handleErr(err)

// 	return &simpleServer{
// 		addr:  addr,
// 		proxy: httputil.NewSingleHostReverseProxy(serverUrl),
// 	}
// }

// type LoadBalancer struct {
// 	port            string
// 	roundRobinCount int
// 	servers         []Server
// }

// func handleErr(err error) {
// 	if err != nil {
// 		fmt.Printf(" error: %+v\n", err)
// 		os.Exit(1)
// 	}
// }

// func (s *simpleServer) Address() string { return s.addr }
// func (s *simpleServer) IsAlive() bool   { return true }
// func (s *simpleServer) Server(rw http.ResponseWriter, req *http.Request) {
// 	s.proxy.ServeHTTP(rw, req)
// }

// func (lb LoadBalancer) getNextAvailableServer() Server {
// 	server := lb.servers[lb.roundRobinCount%len(lb.servers)]

// 	for !server.IsAlive() {
// 		lb.roundRobinCount++
// 		server = lb.servers[lb.roundRobinCount%len(lb.servers)]
// 	}
// 	lb.roundRobinCount++
// 	return server
// }

// func (lb *LoadBalancer) serveProxy(rw http.ResponseWriter, req *http.Request) {

// 	targetServer := lb.getNextAvailableServer()
// 	fmt.Printf("fowarding request to address %q\n", targetServer.Address())
// 	targetServer.Server(rw, req)
// }

// func main() {

// 	servers := []Server{
// 		newSimpleServer("https://www.bing.com/"),
// 		newSimpleServer("https://www.facebook.com/"),
// 		newSimpleServer("https://duckduckgo.com/"),
// 	}

// 	lb := NewLoadBalancer("8000", servers)
// 	handleRedirect := func(rw http.ResponseWriter, req *http.Request) {
// 		lb.serveProxy(rw, req)
// 	}
// 	http.HandleFunc("/", handleRedirect)
// 	fmt.Printf("serving request at localhost:%s\n", lb.port)
// 	http.ListenAndServe(":"+lb.port, nil)

// }
