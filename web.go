package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// Log the HTTP request
func logHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
}

// Write a "Powered by $POWERED_BY" response using the environment variable.
func poweredByHandler(w http.ResponseWriter, r *http.Request) {
	poweredBy := os.Getenv("POWERED_BY")
	if poweredBy == "" {
		poweredBy = "Deis"
	}
	// Print the string to the ResponseWriter
	fmt.Fprintf(w, "Powered by %v\n", poweredBy)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	logHandler(w, r)
	poweredByHandler(w, r)
}

// Start an HTTP server which dispatches handlers based on URL
func main() {
	http.HandleFunc("/", rootHandler)
	port := os.Getenv("PORT")
	log.Printf("server is listening on %v...\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}
