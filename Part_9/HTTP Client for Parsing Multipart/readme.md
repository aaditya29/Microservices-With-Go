## HTTP Client Multipart Form-Data Request

The example written in [code](https://github.com/aaditya29/Microservices-With-Go/blob/master/Part_9/HTTP%20Client%20for%20Parsing%20Multipart/client.go) sends the multipart/form-data request body in an HTTP request to the server created in this [example](https://github.com/aaditya29/Microservices-With-Go/blob/master/Part_9/HTTP%20Server%20For%20Sending%20Multipart/server.go) <br>

Here, first we have to create a [multipart-writer](https://golang.org/pkg/mime/multipart/#Writer)

	writer := multipart.NewWriter(body)

This multipart writer will provide two methods:
a) CreateFormField  - Used to create a text field to be sent in the multipart request body. We will create the name and age<br>
field using this method. <br>
b) CreateFormFile - Used to create a file field to be sent in the multipart request body. We will create the photo field <br>
using this method.<br>

For more detailed explanation of program move to [code](https://github.com/aaditya29/Microservices-With-Go/blob/master/Part_9/HTTP%20Client%20for%20Parsing%20Multipart/client.go) <br>


