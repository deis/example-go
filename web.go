package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// Start an HTTP server listening on $PORT which dispatches to a rootHandler.
func main() {
	http.HandleFunc("/", rootHandler)
	port := os.Getenv("PORT")
	log.Printf("listening on %v...\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}

// Return a "Powered by $POWERED_BY" message using the environment variable.
func poweredByHandler(w http.ResponseWriter, r *http.Request) {
	release := os.Getenv("DEIS_RELEASE")
	if release == "" {
		release = "<unknown>"
	}
	powered := os.Getenv("POWERED_BY")
	if powered == "" {
		powered = "Deis"
	}
	// Print the string to the ResponseWriter
	hostname, _ := os.Hostname()
	fmt.Fprintf(w, "Powered by %v\nRelease %v on %v\n", powered, release, hostname)
}

// Log the HTTP request
func logHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	logHandler(w, r)
	poweredByHandler(w, r)
}
