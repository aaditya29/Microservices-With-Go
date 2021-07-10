## Creating RPC Client

A program that sends requests to the RPC server and processes the response is an RPC client. Normally, an RPC client sends<br>
a POST HTTP request (with a procedure payload in the body) to the server but depending on the server implementation, a <br> client can implement a different mode of communication.<br>

In Golang, an RPC server uses a standard HTTP server on a specific endpoint. An RPC client must connect to this endpoint <br> using CONNECT method. This approach is called [HTTP Tunneling](https://en.wikipedia.org/wiki/HTTP_tunnel). Here, the HTTP server acts as a proxy server and facilitates the communication between RPC client and server. <br>

The [logic](https://golang.org/src/net/rpc/server.go?s=21597:21656#L700) of connecting with a client is already implemented by the RPC server using <b*Server.HandleHTTP</b> method.<br>
We just need a client that can communicate with this server using an appropriate mechanism and it should be able to handle the responses. <br>

The rpc package also provides the Client structure type that defines a self-sufficient RPC client designed to communicate with the Go’s RPC server. <br>
The rpc.DialHTTP function establishes a connection with the RPC server hosted ussing an HTTP server at DefaultRPCPath endpoint. This function returns a *Client object which will be our RPC client.<br>

	func DialHTTP(network, address string) (*Client, error)

Here, the network argument specifies the transport protocol. Since we need to connect with an RPC server using HTTP Tunneling which operates on TCP, this argument value will be tcp. <br>
The address argument is the physical address of the server on the network along with the server port number. <br>

The implementation of creating the client can be found [here](https://github.com/aaditya29/Microservices-With-Go/blob/master/Part_10/Creating%20An%20RPC%20Client/creatingclient.go).

