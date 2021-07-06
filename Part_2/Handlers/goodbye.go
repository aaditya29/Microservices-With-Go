package handlers

import (
	"log"
	"net/http"
)

// Making a handler goodbye
type Goodbye struct {
	l *log.Logger
}

// NewHGoodbye creates a new goodbye handler with the given logger
func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

// ServeHTTP implements the go http.Handler interface
//We can read about it on:  https://golang.org/pkg/net/http/#Handler
func (h *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Goodbye"))
}
