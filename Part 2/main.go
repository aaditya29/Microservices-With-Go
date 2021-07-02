package main

import (
	"log"
	"net/http"
	"os"
	"time"
	"github.com/aaditya29/Microservices-With-Go/Part 2/Handlers/"
)

func main() {
	l := log.New(os.Stdout, "products-api ", log.LstdFlags) //Creating new logger
	//Creating a reference to the handler
	hh := handlers.NewHello(l)   //HelloHandler
	gh := handlers.NewGoodbye(l) //GoodbyeHandler

	// creating a new serve mux and registering the handlers
	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)

	//Creating a new server
	// A server defines parameters for running an HTTP server.
	//The zero value for server is a valid configuration.
	s := http.Server{
		Addr:         ":9090",           // configure the bind address
		Handler:      sm,                // set the default handler
		ErrorLog:     l,                 // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}
	s.ListenAndServe()
}
