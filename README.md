# go-pprof-hijack-middleware

pprof-hijack-middleware is a `net/http` middleware that gathers pprof cpu/mem profiles for the lifecycle of an HTTP
request.

It is mostly intended to be used for development purposes and is useful to limit the pprof sample to
the exact lifecycle of an HTTP request handler.

It hijacks the original `http.ResponseWriter` and returns the pprof profile data (gzipped compressed) as the HTTP response payload.

The HTTP request must generate enough load/allocs for the pprof profiling rate to pick up.

A simple "hello world" handler will, most likely, not output any useful profiling data.

There are no dependencies besides the standard lib.

### Install with go get:

```
go get github.com/miguelcnf/go-pprof-hijack-middleware
```

### Use as any regular `net/http` middleware:

```go
mux := http.NewServeMux()

cpuProfileHijackHandler := pprofhijackmiddleware.CPUProfile(http.HandlerFunc(customHandler))
mux.Handle("/custom", cpuProfileHijackHandler)

err := http.ListenAndServe(":8080", mux)
if err != http.ErrServerClosed {
    log.Fatalf("server closed unexpectedly: %v", err)
}
```

See the [CPU Profile Example](_examples/README.md) for a running example.