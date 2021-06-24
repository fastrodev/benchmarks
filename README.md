# Benchmarks

Benchmark results depend on your machine. To generate in your local machine, you
must install [`wrk`](https://github.com/wg/wrk), clone, and setup this repo.

```
git clone git@github.com:fastrodev/benchmarks.git
cd benchmarks
make benchmark
```

## Fastrex

Source code: [`fastrex/main.go`](fastrex/main.go)\
Command: `wrk http://localhost:9000`\
Results:

```
Running 10s test @ http://localhost:9000
  2 threads and 10 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    88.76us   93.19us   6.64ms   97.45%
    Req/Sec    48.11k     3.78k   62.87k    82.09%
  961962 requests in 10.10s, 111.01MB read
Requests/sec:  95249.11
Transfer/sec:     10.99MB

```

## Standard net/http

Source code: [`standard/main.go`](standard/main.go)\
Command: `wrk http://localhost:9001`\
Results:

```
Running 10s test @ http://localhost:9001
  2 threads and 10 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    92.94us   40.07us   1.10ms   90.95%
    Req/Sec    44.56k     2.01k   54.48k    75.00%
  887011 requests in 10.00s, 108.28MB read
Requests/sec:  88700.49
Transfer/sec:     10.83MB

```

## Standard Node.js

Source code: [`node/main.js`](node/main.js)\
Command: `wrk http://localhost:3001`\
Results:

```
Running 10s test @ http://localhost:3001
  2 threads and 10 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   208.46us  259.72us  13.17ms   98.21%
    Req/Sec    25.48k     3.53k   29.58k    86.63%
  512034 requests in 10.10s, 65.43MB read
Requests/sec:  50696.05
Transfer/sec:      6.48MB

```

## Express.js

Source code: [`express/main.js`](express/main.js)\
Command: `wrk http://localhost:3000`\
Results:

```
Running 10s test @ http://localhost:3000
  2 threads and 10 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     1.16ms  672.19us  21.38ms   92.83%
    Req/Sec     4.53k   809.90     5.42k    84.16%
  90991 requests in 10.10s, 20.74MB read
Requests/sec:   9006.68
Transfer/sec:      2.05MB

```
