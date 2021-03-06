// pprofhijackmiddleware provides pprof cpu/mem profile gathering as a net/http middleware.
package pprofhijackmiddleware

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"runtime"
	"runtime/pprof"
	"strconv"
)

const hijackRecordedHTTPCode = "pprof-hijack-middleware-recorded-http-code"

// CPUProfile gathers a pprof CPU profile during the lifecycle of the passed HTTP request handler.
// It hijacks the http.ResponseWriter to return a pprof compatible gzipped file as the HTTP response body.
func CPUProfile(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		buffer := bytes.Buffer{}
		err := pprof.StartCPUProfile(&buffer)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		recorder := httptest.NewRecorder()
		h.ServeHTTP(recorder, r)
		w.Header().Set(hijackRecordedHTTPCode, strconv.Itoa(recorder.Code))

		pprof.StopCPUProfile()

		_, _ = w.Write(buffer.Bytes())
	})
}

// MemProfile gathers a pprof Heap profile after the HTTP request handler is executed and a forced GC is executed.
// It hijacks the http.ResponseWriter to return a pprof compatible gzipped file as the HTTP response body.
func MemProfile(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		recorder := httptest.NewRecorder()
		h.ServeHTTP(recorder, r)
		w.Header().Set(hijackRecordedHTTPCode, strconv.Itoa(recorder.Code))

		runtime.GC()

		buffer := bytes.Buffer{}
		err := pprof.WriteHeapProfile(&buffer)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		_, _ = w.Write(buffer.Bytes())
	})
}
