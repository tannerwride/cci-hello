package main

import (
	"log"
	"net/http"
)

// The handler function
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello. This is a site to demo remote docker"))
}

func main() {
	// Router. Mapping between url and handler function
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	// The server
	log.Println("Starting Server")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
