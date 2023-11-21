package main

import (
	"net/http"
)

func basicHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET requests are allowed", http.StatusMethodNotAllowed)
		return
	}
	response := []byte("Hello, world!")
	w.Write(response)
}

func main() {
	http.HandleFunc("/", basicHandler)
	port := ":8080"
	http.ListenAndServe(port, nil)
}
