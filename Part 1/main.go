package main

import (
	"log"
	"net/http"
)

func main() {
	//Creating fucntion handler which prints hello world on any path request
	http.HandleFunc("/", func(http.ResponseWriter, *http.Request) {
		log.Println("Hello World")
	})

	//Creating function handler which prints goodbye world if the requested path is goodbye
	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Goodbye World")
	})

	http.ListenAndServe("9090", nil) //Binding all the addresses at port 9090
}
