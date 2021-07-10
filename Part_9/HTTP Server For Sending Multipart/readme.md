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

It has two parts. The Value holds all the non-files data. So it will hold the name and age key data.  The File part holds<br>
all file data. So it will hold the data for photo key. <br>
Both the parts have their value part as an array because for the same key there can be multiple values. <br>

b) Form: It holds combine data of query string and non-file fields of the multipart/form-data request body. <br>
For example, for the above case, it will only hold the name and age field. It has below format:<br>

	map[string][]string

c) PostForm: It only holds the non-file fields of the multipart/form-data request body. So for the above case, it will only<br> hold the name and age field. It also has the same format as <b>Form.</b>

	map[string][]string

Other than the above fields, the request object also provides the below utility function: <br>
d) FormFile– This function can be used to get the first file for the given key. So this function will only return data <br>
for photo key and not for name and age key.

e) PostFormValue – This function returns the first value for the given form key.  This function will only return data for<br>
name and age key and not for photo key.

#### Accessing Name, Key and Photo Values in Different Way:

	request.Form["name"]
    request.PostForm["name"] 
    request.MultipartForm.Value["name"]

is going to return

	["John"]

It is going to return array because for the same key there can be different values.

Whereas

	request.PostFormValue("name") 

will return

	John

As we can see it is not an array because as mentioned above this function returns the first value associated with the key. <br>