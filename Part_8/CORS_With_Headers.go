package main

import "net/http"

/*
Headers for CORS:
The following are the new HTTP headers added by the CORS standard:
1. Access-Control-Allow-Origin
2. Access-Control-Allow-Credentials
3. Access-Control-Allow-Headers
4. Access-Control-Allow-Methods etc.

In Golang, we need to add CORS headers in OPTIONS method of HTTP request.
*/

// We need to create a function which will add all CORS headers

func setupCorsResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
}

// And then we will have to create HTTP Handlers

// Register routing for creating customer
route.Handle("/"+APIVersion+"/customers", createUser())
   Methods(http.MethodPost, http.MethodOptions)
   Name("Create customer")

// Calling function to add all CORS Headers
func createUser(w http.ResponseWriter, req *http.Request) {
	setupCorsResponse(&w, req)
	if (*req).Method == "OPTIONS" {
	   return
	}
 
	// processing the request...
 }

 // After adding this, server will send CORS header in response and browser will accept it
 //and will not restrict any content of this server.
