## Creating RPC Server

Go provides <b>net/rpc</b> package to implement RPC client-server architecture. This package exposes methods and types to <br> create RPC servers to host RPC services and RPC clients to connect with these RPC services. <br>

### Procedure to Create RPC Server

First we need some procedures so that we register them with the server and clients can execute them remotely. <br>
In Go, we do not register functions or methods directly. We register a type that implements methods. <br>
