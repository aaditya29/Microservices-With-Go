## Creating RPC Client

A program that sends requests to the RPC server and processes the response is an RPC client. Normally, an RPC client sends<br>
a POST HTTP request (with a procedure payload in the body) to the server but depending on the server implementation, a <br> client can implement a different mode of communication.<br>

In Golang, an RPC server uses a standard HTTP server on a specific endpoint. An RPC client must connect to this endpoint <br> using CONNECT method. This approach is called [HTTP Tunneling](https://en.wikipedia.org/wiki/HTTP_tunnel). Here, the HTTP server acts as a proxy server and facilitates the communication between RPC client and server. <br>
