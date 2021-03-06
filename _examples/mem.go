package main

import (
	"log"
	"math/rand"
	"net/http"

	"github.com/miguelcnf/go-pprof-hijack-middleware"
)

var data []float64

var memWorkHandler = func(w http.ResponseWriter, r *http.Request) {
	data = make([]float64, 10000000)
	for i := range data {
		data[i] = rand.Float64()
	}
	w.WriteHeader(http.StatusOK)
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/mem", http.HandlerFunc(memWorkHandler))

	err := http.ListenAndServe(":8080", pprofhijackmiddleware.MemProfile(mux))
	if err != http.ErrServerClosed {
		log.Fatalf("server closed unexpectedly: %v", err)
	}
}
