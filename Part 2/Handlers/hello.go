package handlers

import "net/http"

//A handler responds to an HTTP request.
//To create a handler we have to create a struct which implements interface http handler.

//Hello is a simple handler
type Hello struct {
}

//Method which satisfies HTTP handler interface
func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) { //Signature that we need to satisfy the HTTP Interface

}
