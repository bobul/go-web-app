package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Create a handler that writes a slice of bytes contains text "Hello from App" as response body.
func home(w http.ResponseWriter, r *http.Request) {
	// Handle only root path, but do not handle other arbitrary paths
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from App"))
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Showing the note with ID: %d...", id)
}
func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method is forbidden!", 405)
		return
	}
	w.Write([]byte("Creating a snippet..."))
}
func main() {
	// Use function http.NewServeMux() for new rooter initialization, then function home acts as handler for url
	// template "/".
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	// Use function http.ListenAndServe for running web server.
	// We pass two argument: TCP network address and created rooter.
	// Then error handling.
	log.Println("Running web server on localhost:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
