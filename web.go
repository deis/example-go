package main

import (
	"fmt"
	"net/http"
	"os"
)

// Start an HTTP server listening on $PORT which dispatches to the
// powered_by() function.
func main() {
	http.HandleFunc("/", powered_by)
	port := os.Getenv("PORT")
	fmt.Printf("listening on %v...\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}

// Return a "Powered by $POWERED_BY" message using the environment variable.
func powered_by(res http.ResponseWriter, req *http.Request) {
	deis_release := os.Getenv("DEIS_RELEASE")
	if deis_release == "" {
		deis_release = "<unknown>"
	}
	powered_by := os.Getenv("POWERED_BY")
	if powered_by == "" {
		powered_by = "Deis"
	}
	// Print the string to the ResponseWriter
	hostname, _ := os.Hostname()
	fmt.Fprintf(res, "Release %v Powered by %v on %v\n", deis_release, powered_by, hostname)
}
