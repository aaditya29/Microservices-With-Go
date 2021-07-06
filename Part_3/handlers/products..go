package handlers

import (
	"log"
	"net/http"

	"github.com/aaditya29/Microservices-With-Go/tree/master/Part-Three/data"
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
	// handle the request for a list of products
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	// catch all
	// if no method is satisfied return an error
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

// getProducts returns the products from the data store
func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	//we want to take getrequest and return product list from folder data
	// we need encoding json for this stuff
	lp := data.GetProducts()
	err := lp.ToJSON(rw) //returning responsewriter
	if err != nil {
		http.Error(rw, "Unable to marshal json or fetch data", http.StatusInternalServerError)
	}
}
