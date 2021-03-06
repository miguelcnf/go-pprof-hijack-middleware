package main

import (
	"crypto/sha512"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	pprofhijackmiddleware "github.com/miguelcnf/go-pprof-hijack-middleware"
)

var cpuWorkHandler = func(w http.ResponseWriter, r *http.Request) {
	hasher := sha512.New()
	random := strconv.FormatInt(rand.Int63(), 10)
	for i := 0; i < 100000000; i++ {
		_, _ = hasher.Write([]byte(random))
	}
	w.WriteHeader(http.StatusOK)
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/cpu", http.HandlerFunc(cpuWorkHandler))

	err := http.ListenAndServe(":8080", pprofhijackmiddleware.CPUProfile(mux))
	if err != http.ErrServerClosed {
		log.Fatalf("server closed unexpectedly: %v", err)
	}
}
