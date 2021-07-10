## HTTP Client Multipart Form-Data Request

The example written in [code](https://github.com/aaditya29/Microservices-With-Go/blob/master/Part_9/HTTP%20Client%20for%20Parsing%20Multipart/client.go) sends the multipart/form-data request body in an HTTP request to the server created in this [example](https://github.com/aaditya29/Microservices-With-Go/blob/master/Part_9/HTTP%20Server%20For%20Sending%20Multipart/server.go) <br>

Here, first we have to create a [multipart-writer](https://golang.org/pkg/mime/multipart/#Writer)