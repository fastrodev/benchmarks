# Benchmarks

Router framework vs Standard net/http
- OS: macOS Big Sur v11.3.1
- Processor: 2 GHz Quad-Core Intel Core i5
- Memory: 16 GB 3733 MHz LPDDR4X
- Command: `wrk http://localhost:9000`

### Router Framework
- Source code: [router/main.go](router/main.go)  	
- Results:
```
Running 10s test @ http://localhost:9000
  2 threads and 10 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    94.56us   45.91us   1.55ms   92.98%
    Req/Sec    44.26k     3.25k   53.75k    89.11%
  889235 requests in 10.10s, 102.61MB read
Requests/sec:  88044.39
Transfer/sec:     10.16MB
```
### Standard net/http 
- Source code: [standard/main.go](standard/main.go)  
- Results:
```
Running 10s test @ http://localhost:9000
  2 threads and 10 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    99.31us  254.38us  12.51ms   99.44%
    Req/Sec    44.97k     1.68k   47.91k    65.35%
  903706 requests in 10.10s, 110.32MB read
Requests/sec:  89475.41
Transfer/sec:     10.92MB
```
### Express.js
- Source code: [express/main.js](express/main.js)  
- Results:
```
Running 10s test @ http://localhost:9000
  2 threads and 10 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.08ms  663.45us  23.39ms   94.09%
    Req/Sec     4.87k   706.79     5.43k    91.09%
  97830 requests in 10.10s, 22.30MB read
Requests/sec:   9685.14
Transfer/sec:      2.21MB
```

### Standard Node.js
- Source code: [node/main.js](node/main.js) 
- Results:
```
Running 10s test @ http://localhost:9000
  2 threads and 10 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   291.17us    1.25ms  35.60ms   99.24%
    Req/Sec    24.59k     4.48k   49.85k    94.53%
  491852 requests in 10.10s, 62.85MB read
Requests/sec:  48699.82
Transfer/sec:      6.22MB
```
