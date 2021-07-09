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
But if you need to upload files then <b>application/x-www-form-urlencoded</b> is not much efficient.
