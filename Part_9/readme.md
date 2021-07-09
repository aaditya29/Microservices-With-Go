## Building Microservices With Golang

#### Part Nine
In this part we are going to learn at how to handle multi-part HTTP requests. <br>
While the modern approach is to use a RESTful endpoints. Multi-part can still be useful, simple HTML browser forms implement<br>
this standard for file uploads and forms.

### HTTP Multipart Overview:
In HTTP context, the multipart/form-data content-type is used for submitting HTML form. In the case of multipart/form-data,<br> as the name suggests, the body consists of different parts separated by a delimiter or boundary where each part is <br>
described by its own headers. The delimiter or boundary is also sent as part of the header only.

When you sending an HTML form through a browser in an HTTP call, the data contents can be sent in as request body in the below two formats: <br>
1. application/x-www-form-urlencoded
2. multipart/form-data

For most of the cases, <b>application/x-www-form-urlencoded</b> can be used.<br>
But if you need to upload files then <b>application/x-www-form-urlencoded</b> is not much efficient.<br>

For example, let’s say following two data are needed to send:
1. name
2. age
Then <b>application/x-www-form-urlencoded</b> can be used to send the above data.
But let’s say that you also need to send the profile photo of the user in the request as well. So the data is now as below:
1. name
2. age
3. photo
In the above case, it will not be efficient to use <b>application/x-www-form-urlencoded</b> content-type.<br>
<b>multipart/form-data</b> should be used in this case. So for sending simple form data use <b>application/x-www-form-urlencoded </b> <br>
but if the form-data also needs to send binary data then use <b>multipart/form-data.</b>


