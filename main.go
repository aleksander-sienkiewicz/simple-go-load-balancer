package main

import (
	"fmt" //print to terminal
	"net/http"
	"net/http/httputil" //import httputil
	"net/url"
	"os" //for error handling
)

type Server interface {
	// Address returns the address with which to access the server
	Address() string //defined adr,isalive,serve in interface.

	// IsAlive returns true if the server is alive and able to serve requests
	IsAlive() bool

	// Serve uses this server to process the request
	Serve(rw http.ResponseWriter, req *http.Request) //takes response writer and request
}

/*
will have adrress, type string
proxy type httpUtillity
*/
type simpleServer struct {
	addr  string
	proxy *httputil.ReverseProxy
}

// defined adr,isalive,serve in interface.
func (s *simpleServer) Address() string { return s.addr }

func (s *simpleServer) IsAlive() bool { return true } //define these three methods

func (s *simpleServer) Serve(rw http.ResponseWriter, req *http.Request) {
	s.proxy.ServeHTTP(rw, req) //proxy is simple server, has adr and proxy(which we get from httpUtil pkg)
} //once we have accress to proxy, since we get from pkg, we have access to serve http
//serve same website just serving thru reverse proxy

/*
Init a new simple server
will have and adress and a proxy
*/
func newSimpleServer(addr string) *simpleServer {
	serverUrl, err := url.Parse(addr) //
	handleErr(err)

	return &simpleServer{ //return a new instance of a server
		addr:  addr,                                          //adr == adr, what we recieve
		proxy: httputil.NewSingleHostReverseProxy(serverUrl), //proxy required to make a new instance of simple server
	}
}

/*
 */
type LoadBalancer struct {
	port            string   //has port
	roundRobinCount int      //has roundrobin count
	servers         []Server //has servers, a slice or a collection of the server we create
}

/*
returns a newload balancer struct
takes in port, servers, what we need to create it
*/
func NewLoadBalancer(port string, servers []Server) *LoadBalancer {
	return &LoadBalancer{
		port:            port,    //port = to port duh
		roundRobinCount: 0,       //0 in beggining
		servers:         servers, //servers =2 servers
	}
}

// handleErr prints the error and exits the program
// Note: this is not how one would want to handle errors in production, but
// serves well for demonstration purposes.
func handleErr(err error) { //recieve error
	if err != nil {
		fmt.Printf("error: %v\n", err) //print error, now we dont have to write this every time
		os.Exit(1)
	}
}

/*
	getNextServerAddr method returns the address of the next available server to send a

request to, using a simple round-robin algorithm
*/
func (lb *LoadBalancer) getNextAvailableServer() Server {
	server := lb.servers[lb.roundRobinCount%len(lb.servers)]
	for !server.IsAlive() { //cycle thru live servers
		lb.roundRobinCount++ //up count
		server = lb.servers[lb.roundRobinCount%len(lb.servers)]
	}
	lb.roundRobinCount++

	return server //return avaibable server
}

/*
takes res. writer and request
method accessible from loadbalancer
*/
func (lb *LoadBalancer) serveProxy(rw http.ResponseWriter, req *http.Request) {
	targetServer := lb.getNextAvailableServer() //assign to target server
	//check avaibality from isAlive

	// could optionally log stuff about the request here!
	fmt.Printf("forwarding request to address %q\n", targetServer.Address())
	//print this to terminal

	// could delete pre-existing X-Forwarded-For header to prevent IP spoofing
	targetServer.Serve(rw, req)

	//IP spoofing is the creation of Internet Protocol (IP) packets which have a modified source
	//address in order to either hide the identity of the sender, to impersonate another computer system, or both.
}

/*
program start
*/
func main() {
	servers := []Server{ //create list of servers using newsimpleserver function
		newSimpleServer("https://www.facebook.com"),
		newSimpleServer("https://www.bing.com"),
		newSimpleServer("https://www.duckduckgo.com"),
	}

	lb := NewLoadBalancer("8000", servers) //start lb on port 8000, pass servers
	//access to these loadbalancers thru 'lb', which has two methods

	handleRedirect := func(rw http.ResponseWriter, req *http.Request) { //called when someone hits root route
		lb.serveProxy(rw, req)
	}

	// register a proxy handler to handle all requests
	http.HandleFunc("/", handleRedirect) //if root route hit redirect :D

	fmt.Printf("serving requests at 'localhost:%s'\n", lb.port) //optional
	//just print this to let us know program started
	http.ListenAndServe(":"+lb.port, nil) //use http package to start our server
	//lb.port represents 8000 we input as our port
}
