package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	l := log.New(os.Stdout, "products-api ", log.LstdFlags) //Creating new logger
	//Creating a reference to the handler
	hh := handlers.NewHello(l)

	// creating a new serve mux and registering the handlers
	sm := http.NewServeMux()
	sm.Handle("/", hh)

	http.ListenAndServe(":9090", nil)
}
