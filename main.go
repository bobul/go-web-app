package main

import (
	"log"
	"net/http"
)

// Create a handler that writes a slice of bytes contains text "Hello from App" as response body.
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from App"))
}
func main() {
	// Use function http.NewServeMux() for new rooter initialization, then function home acts as handler for url
	// template "/".
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	// Use function http.ListenAndServe for running web server.
	// We pass two argument: TCP network address and created rooter.
	// Then error handling.
	log.Println("Running web server on localhost:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
