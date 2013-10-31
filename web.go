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
	powered_by := os.Getenv("POWERED_BY")
	if powered_by == "" {
		powered_by = "Deis"
	}
	// Print the string to the ResponseWriter
	fmt.Fprintf(res, "Powered by %v\n", powered_by)
}
