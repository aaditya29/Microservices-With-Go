## HTTP Server Example

For HTTP server an incoming request is represented by the [request](https://golang.org/src/net/http/request.go) struct. <br>

To parse the <b>multipart/form-data</b> request body we need to first call the below function on the request object:<br>

	request.ParseMultipartForm()


