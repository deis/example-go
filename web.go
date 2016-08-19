package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// main starts an HTTP server listening on $PORT which dispatches to request handlers.
func main() {
	http.Handle("/healthz", http.HandlerFunc(healthcheckHandler))
	// wrap the poweredByHandler with logging middleware
	http.Handle("/", logRequestMiddleware(http.HandlerFunc(poweredByHandler)))
	port := os.Getenv("PORT")
	log.Printf("listening on %v...\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}

// poweredByHandler writes "Powered by $POWERED_BY" to the response.
func poweredByHandler(w http.ResponseWriter, r *http.Request) {
	release := os.Getenv("WORKFLOW_RELEASE")
	if release == "" {
		release = "<unknown>"
	}
	powered := os.Getenv("POWERED_BY")
	if powered == "" {
		powered = "Deis"
	}
	// print the string to the ResponseWriter
	hostname, _ := os.Hostname()
	fmt.Fprintf(w, "Powered by %v\nRelease %v on %v\n", powered, release, hostname)
}

// healthcheckHandler returns 200 for kubernetes healthchecks.
func healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte{})
}

// logRequestMiddleware writes out HTTP request information before passing to the next handler.
func logRequestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		remote := r.RemoteAddr
		if forwardedFor := r.Header.Get("X-Forwarded-For"); forwardedFor != "" {
			remote = forwardedFor
		}
		log.Printf("%s %s %s", remote, r.Method, r.URL)
		// pass the request to the next handler
		next.ServeHTTP(w, r)
	})
}
