package main

import (
	"log"
	"net/http"
)

func main() {
	// Signin and Welcome are the handlers that will be implemented in the server
	http.HandleFunc("/signin", Signin)
	http.HandleFunc("/welcome", Welcome)
	http.HandleFunc("/refresh", Refresh)

	log.Fatal(http.ListenAndServe(":8000", nil)) // start the server on port 8000
}
