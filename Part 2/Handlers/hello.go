package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//A handler responds to an HTTP request.
//To create a handler we have to create a struct which implements interface http handler.

//Hello is a simple handler
type Hello struct {
	l *log.Logger
}

// NewHello creates a new hello handler with the given logger
// It is going to take log.logger and going to return Hello handler
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

//Method which satisfies HTTP handler interface
func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) { //Signature that we need to satisfy the HTTP Interface
	log.Println("Hello World")

	//Reading the message from the user or reading the body
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Oops! Error Occured", http.StatusBadRequest)
		return
	}

	//Writing The Response
	fmt.Fprintf(rw, "Hello %s", d)
}
