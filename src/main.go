package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 5000

// Handler for the root path
func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Request received on path: %s\n", r.URL.Path)
}

// Main function to start the HTTP server
func main() {
	http.HandleFunc("/", index)
	addr := fmt.Sprintf("0.0.0.0:%d", port)
	log.Printf("Go app version 0.2 running on: %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
