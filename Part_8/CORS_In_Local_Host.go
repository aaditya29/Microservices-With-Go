func handler(w http.ResponseWriter, req *http.Request) {
	// ...
	enableCors(&w)
	// ...
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}