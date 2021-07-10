## Creating RPC Server

Go provides <b>net/rpc</b> package to implement RPC client-server architecture. This package exposes methods and types to <br> create RPC servers to host RPC services and RPC clients to connect with these RPC services. <br>

### Procedure to Create RPC Server

First we need some procedures so that we register them with the server and clients can execute them remotely. <br>
In Go, we do not register functions or methods directly. We register a type that implements methods. <br>

To say in a more clear way, we register an instance (object) of that <b>type</b>. All public methods of that type are registered<br>
as procedures automatically. <br>
But there are following conditions that this <b>type</b> should satisfy:
1. This type should be exported, which means it should start with an uppercase letter.<br>
2. The method that needs to be used as a procedure should be exported as well, which means they should start with an uppercase letter.
3. This method should have precisely two arguments and types of these arguments should also be exported (except built-in types).
4. The second argument of this method should be a pointer. This argument will be used to return a result by overriding its value.
5. This method should return an error value. This is useful to return an error in case if the procedure call fails with an error.

Now we will create a common package that exports types that can be shared between RPC clients and servers in this [code](https://github.com/aaditya29/Microservices-With-Go/blob/master/Part_10/Creating%20An%20RPC%20Server/creatingserver.go).<br>
And then in second [part](https://github.com/aaditya29/Microservices-With-Go/blob/master/Part_10/Creating%20An%20RPC%20Server/implementingserver.go) we will implement the server. <br>
