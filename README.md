# Benchmarks

Router framework vs Standard net/http
- OS: macOS Big Sur v11.3.1
- Processor: 2 GHz Quad-Core Intel Core i5
- Memory: 16 GB 3733 MHz LPDDR4X
- Command: `wrk http://localhost:9000`

### Rider Framework
- Source code: [rider/main.go](rider/main.go)  	
- Results:
```
Running 10s test @ http://localhost:9000
  2 threads and 10 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    89.40us   51.49us   1.53ms   92.58%
    Req/Sec    47.19k     4.32k   52.35k    77.72%
  948152 requests in 10.10s, 109.41MB read
Requests/sec:  93878.51
Transfer/sec:     10.83MB
```
### Standard net/http 
- Source code: [standard/main.go](standard/main.go)  
- Results:
```
Running 10s test @ http://localhost:9000
  2 threads and 10 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    94.82us   46.17us   1.53ms   92.53%
    Req/Sec    43.95k     2.14k   54.25k    82.67%
  883422 requests in 10.10s, 107.84MB read
Requests/sec:  87464.59
Transfer/sec:     10.68MB
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
