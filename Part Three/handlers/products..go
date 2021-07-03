package handlers

import (
	"log"
	"net/http"
)

// Products is a http.Handler
type Products struct {
	l *log.Logger
}

// NewProducts creates a products handler with the given logger
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// ServeHTTP is the main entry point for the handler and staisfies the http.Handler
// Making ServeHTTP function
func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

}
