package main

import "github.com/dgrijalva/jwt-go"

// Writing the JWT Key used to create the signature
var jwtKey = []byte("my_secret_key")

// Handling User signin
// It will take the users credentials and log them in.
// Storing the users information as an in-memory map in our code
var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

// Create a struct that models the structure of a user, both in the request body, and in the Database
type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

// Creating a struct that will be encoded to JWT
// Adding jwt.StandardClaims as an embedded type to provide fields like starting time
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
