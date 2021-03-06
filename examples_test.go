package pprofhijackmiddleware_test

import (
	"fmt"
	"log"
	"net/http"

	pprofhijackmiddleware "github.com/miguelcnf/go-pprof-hijack-middleware"
)

var (
	customHandler = func(w http.ResponseWriter, r *http.Request) {}
	helloHandler  = func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "hello")
	}
)

func ExampleCPUProfile() {
	mux := http.NewServeMux()

	// regular handler, not hijacked
	mux.Handle("/hello", http.HandlerFunc(helloHandler))

	// hijack a single handler
	cpuProfileHijackHandler := pprofhijackmiddleware.CPUProfile(http.HandlerFunc(customHandler))
	mux.Handle("/custom", cpuProfileHijackHandler)

	err := http.ListenAndServe(":8080", mux)
	if err != http.ErrServerClosed {
		log.Fatalf("server closed unexpectedly: %v", err)
	}
}

func ExampleMemProfile() {
	mux := http.NewServeMux()
	mux.Handle("/hello", http.HandlerFunc(helloHandler))
	mux.Handle("/custom", http.HandlerFunc(customHandler))

	// hijack all registered handlers
	err := http.ListenAndServe(":8080", pprofhijackmiddleware.MemProfile(mux))
	if err != http.ErrServerClosed {
		log.Fatalf("server closed unexpectedly: %v", err)
	}
}
