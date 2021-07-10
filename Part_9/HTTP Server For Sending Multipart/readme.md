## HTTP Server Example

For HTTP server an incoming request is represented by the [request](https://golang.org/src/net/http/request.go) struct. <br>

To parse the <b>multipart/form-data</b> request body we need to first call the below function on the request object:<br>

	request.ParseMultipartForm()

What above function will do is that it will parse the incoming request body and the data will be loaded in the below fields <br>
of the request object:

a) Multipart Form: The entire multipart/form-data request body will be loaded into this field. For example, in the example <br>
written in home page readme, it will hold the name, age field as well as photo field. It is represented in the below format:

	type Form struct {
	Value map[string][]string
	File  map[string][]*FileHeader
    }

