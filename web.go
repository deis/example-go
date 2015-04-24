package main

import (
	"fmt"
	"net/http"
	"os"
)

// Start an HTTP server listening on $PORT which dispatches to the
// poweredBy() function.
func main() {
	http.HandleFunc("/", poweredBy)
	port := os.Getenv("PORT")
	fmt.Printf("listening on %v...\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}

// Return a "Powered by $POWERED_BY" message using the environment variable.
func poweredBy(res http.ResponseWriter, req *http.Request) {
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
	fmt.Fprintf(res, "Powered by %v\nRelease %v on %v\n", powered, release, hostname)
}
