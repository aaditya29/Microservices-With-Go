## JWT Authentication

<b>JSON Web Token (JWT)</b> is a JSON encoded representation of a claim(s) that can be transferred between two parties. The claim is digitally signed by the issuer of the token, and the party receiving this token can later use this digital signature to prove the ownership on the claim. <br>

JWTs can be broken down into three parts:
1. Header
2. Payload
3. ignature.

#### Header
The information contained in the header describes the algorithm used to generate the signature. The decoded version of the header from the above example looks like:<br>


	{ “alg”: “HS256”,
    “typ”: “JWT”
    }

[HS256](https://www.loginradius.com/blog/async/jwt-signing-algorithms/) is the hashing algorithm HMAC SHA-256 used to generate the signature in the above code.<br>


