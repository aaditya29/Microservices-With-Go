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
	
	// starting the server
	go func() {
		l.Println("Starting server on port 9090")

		err := s.ListenAndServe()//Handling listenandserve in gofunc so that it not gonna block
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)// Going to broadcast a message whenever os kills a command or interrupts which is written on previous line cod

	// Block until a signal is received.
	sig := <-c//message availlable is consumed
	log.Println("Got signal:", sig)

	// When message is consumed gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)//gracefull shutdown takes context as parameter

}
