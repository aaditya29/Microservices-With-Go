// Building Restful services with GOlang
// REST stands for Representational State Transfer Services
// It is an architectural approach to design web services.

package main

import "net/http"

func main() {
	// create the handlers
	ph := handlers.NewProducts(l)

	// create a new serve mux and register the handlers
	sm := http.NewServeMux()
	sm.Handle("/", ph)
}
