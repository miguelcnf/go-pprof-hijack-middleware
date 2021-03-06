### CPU Profile

Run the server:
```
$ go run _examples/cpu.go 
...
```

Call the hijacked HTTP handler:

```
$ curl 0:8080/cpu -s --output cpuprofile.pb.gz 
```

Load the profile with go tool:

```
$ go tool pprof cpuprofile.pb.gz 
Type: cpu
Time: Mar 6, 2021 at 3:20pm (WET)
Duration: 7.53s, Total samples = 7.16s (95.14%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 6030ms, 84.22% of 7160ms total
Dropped 62 nodes (cum <= 35.80ms)
Showing top 10 nodes out of 72
      flat  flat%   sum%        cum   cum%
    2120ms 29.61% 29.61%     2120ms 29.61%  crypto/sha512.blockAVX2
    1680ms 23.46% 53.07%     1680ms 23.46%  runtime.kevent
     530ms  7.40% 60.47%     1240ms 17.32%  runtime.mallocgc
     430ms  6.01% 66.48%      430ms  6.01%  runtime.memmove
     320ms  4.47% 70.95%      320ms  4.47%  runtime.nextFreeFast
     220ms  3.07% 74.02%     2640ms 36.87%  crypto/sha512.(*digest).Write
     220ms  3.07% 77.09%     1610ms 22.49%  runtime.rawbyteslice
     190ms  2.65% 79.75%      190ms  2.65%  runtime.pthread_cond_wait
     180ms  2.51% 82.26%     1890ms 26.40%  runtime.stringtoslicebyte
     140ms  1.96% 84.22%     4670ms 65.22%  main.glob..func1
(pprof) 
```

Or Load the profile with pprof:

```
$ pprof -http=localhost:9090 cpuprofile.pb.gz 
Serving web UI on http://localhost:9090
```

Explore the profiling data, for example get a flamegraph:

![CPU Flamegraph](cpu-flamegraph.png?raw=true "CPU Flamegraph")
