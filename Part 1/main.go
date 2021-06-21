//Writing Basic Web Server In GO
package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	//Creating fucntion handler which prints hello world on any path request
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello World")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Oops! Error Occured", http.StatusBadRequest)
			return
		}
	})

	//Creating function handler which prints goodbye world if the requested path is goodbye
	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Goodbye World")
	})

	http.ListenAndServe("9090", nil) //Binding all the addresses at port 9090
}
