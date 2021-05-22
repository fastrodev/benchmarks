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
    Latency    92.64us   41.71us   1.20ms   91.74%
    Req/Sec    44.89k     2.47k   60.14k    88.06%
  897356 requests in 10.10s, 109.54MB read
Requests/sec:  88850.50
Transfer/sec:     10.85MB
```
### Standard net/http 
- Source code: [standard/main.go](standard/main.go)  
- Results:
```
Running 10s test @ http://localhost:9000
  2 threads and 10 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    92.34us   40.24us   1.16ms   90.21%
    Req/Sec    44.70k     2.77k   52.91k    88.61%
  898071 requests in 10.10s, 109.63MB read
Requests/sec:  88920.07
Transfer/sec:     10.85MB
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
