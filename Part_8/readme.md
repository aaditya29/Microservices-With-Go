## Building Microservices With Golang

#### Part Eight
In this part we are going to learn at how to handle Cross <b>Origin Resource Sharing(CORS)</b> in Golang<br>

### CORS
<b>CORS</b> stands for Cross-Origin Resource Sharing, it’s a mechanism that allow browser to access the content of other <br> 
websites on their website. It adds security as it do not allow to access the content without the consent of author.

#### Identification Of Different Domains:
When the domains have:
1. Different name (Like https://example.com and https://demo.com)
2. Different sub domain (Like https://one.example.com and https://two.example.com)
3. Different port (Like https://example.com:8080 and https://example.com:8081)
4. Different protocol (Like https://example.com and http://example.com)

    To see content of different domain, we should keep following points in mind:
1. Either we need to disable security issue from browser so that it doesn’t check for CORS security.
2. Or from server side we need to add some special headers values in HTTP requests and responses so that browser can <br>
understand that other domain is giving full permission to show his content on different domain.

Browser blocks content by CORS Policy and it can be solved by server side by setting special headers.

<b> In Golang, we need to add CORS headers in OPTIONS method of HTTP request.</b>
