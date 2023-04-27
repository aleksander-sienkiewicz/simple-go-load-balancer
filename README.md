# simple go load balancer
 Simple load balancer program made w/ GO 

Load balancers will be implemented as a struct. 3 atributes: port, servers, round robin count
-> function- create a new load balancer

Servers is referenced in loadbalancer struct, itwill be a struct with 2 attributes: Address, Proxy
-> Function create a new server

TWO METHODS:
Get address of a server
Get if the server isAlive

SERVER INTERFACE: (3 methods)
Address()
isAlive()
Serve()


MAIN FUNC(create server list) -> Server Proxy func(make a request or redirect request)-> Get next available server func (check here if server is alive, go to next server use round robin iterator.

SOOOO ya... 
we got our main func, which creates slices of multiple servers. each server represents a dif address, ie. fb.com, bing.com, duckduckgo.com.

We create a load balancer on our specified port, sending the servers to it. then we use our handle func, to handle our /route, so when we hit our local host, we call our the redirect function

redirect func will call the serveProxy method to get next available server.

getNextAvaibaleServer is going to loop between servers and check which are alive, and return a server to our serve func.

Program uses net/http/httputil package for http Utillity 

SEE CLI LOG BELOW, easy set up, easy use. :D

(base) aleksandersienkiewicz@Aleksanders-MacBook-Air ~ % mkdir go-loadbalancer 
(base) aleksandersienkiewicz@Aleksanders-MacBook-Air ~ %           
(base) aleksandersienkiewicz@Aleksanders-MacBook-Air ~ % 
(base) aleksandersienkiewicz@Aleksanders-MacBook-Air ~ % cd documents
(base) aleksandersienkiewicz@Aleksanders-MacBook-Air documents % mkdir go-loadbalancer
(base) aleksandersienkiewicz@Aleksanders-MacBook-Air documents % cd go-loadbalancer
(base) aleksandersienkiewicz@Aleksanders-MacBook-Air go-loadbalancer % cd src
cd: no such file or directory: src
(base) aleksandersienkiewicz@Aleksanders-MacBook-Air go-loadbalancer % mkdir src
(base) aleksandersienkiewicz@Aleksanders-MacBook-Air go-loadbalancer % cd src
(base) aleksandersienkiewicz@Aleksanders-MacBook-Air src % touch main.go
//create files up to this point

//build func. to access launch a private browser window and search http://localhost:8000
(base) aleksandersienkiewicz@Aleksanders-MacBook-Air src % go run main.go
serving requests at 'localhost:8000'
forwarding request to address "https://www.facebook.com"
forwarding request to address "https://www.bing.com"
forwarding request to address "https://www.duckduckgo.com"
 





