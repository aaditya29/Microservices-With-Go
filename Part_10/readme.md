## Remote Procedural Call in Golang

RPC is an inter-process communication architecture to facilitate easy communication between two machines. <br>

### What is Procedure?
A procedure or a subroutine is a collection of instructions in a program to perform a specific task. This could be a function<br>
or a method of a class. When this procedure is invoked on the same machine, such as using a function call syntax <b>producedure()</b><br>
it is called a local procedure call. <br>

When a procedure is executed remotely from a remote machine, it is called a remote procedure call. We can execute a <br>
remote procedure using any type of network connection like HTTP, WebSocket, AMQP, etc. <br>

### RPC Architecture
RPC architecture works in same way as REST architecture. Like REST, RPC is also a client-server protocol where the client<br> sends the server a request to perform some operation and server sends a response.<br>
However, the HTTP method name does not possess any significance but the endpoint name has special significance. <br>

