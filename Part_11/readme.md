## JWT Authentication

<b>JSON Web Token (JWT)</b> is a JSON encoded representation of a claim(s) that can be transferred between two parties. The claim is digitally signed by the issuer of the token, and the party receiving this token can later use this digital signature to prove the ownership on the claim. <br>

JWTs can be broken down into three parts:
1. Header
2. Payload
3. Signature.

#### Header
The information contained in the header describes the algorithm used to generate the signature. The decoded version of the header from the above example looks like:<br>


	{ “alg”: “HS256”,
    “typ”: “JWT”
    }

[HS256](https://www.loginradius.com/blog/async/jwt-signing-algorithms/) is the hashing algorithm HMAC SHA-256 used to generate the signature in the above code.<br>

#### Payload
All the claims within <b>JWT</b> authentication are stored in this part. Claims are used to provide authentication to the party receiving the token.  <br>

For example, a server can set a claim saying ‘isAdmin: true’ and issue it to an administrative user upon successfully logging into the application. The admin user can now send this token in every consequent request he/she sends to the server to prove their identity.<br>

The decoded version of the payload from the JWT example written above looks like:


	{ “sub”: “1234567890”,
    “name”: “John Doe”,
    “iat”: 1516239022
    }

The ‘name’ field written abover is used to identify the user to whom the token was issued to. The ‘sub’ and ‘iat’ are examples of registered claims and are stand for ‘subject’ and ‘issued at’. <br>

#### Signature
The signature part of a JWT is derived from the header and payload fields. The steps involved in creating this signature are as follows:

1.  Combine the base64url encoded representations of header and payload with a dot (.)


	base64UrlEncode(header) + “.” + base64UrlEncode(payload)


