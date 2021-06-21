package main

import "net/http"

func main() {
	http.ListenAndServe("9090", nil) //Binding all the addresses at port 9090
}
