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
{{fastrex}}
```

## Standard net/http

Source code: [`standard/main.go`](standard/main.go)\
Command: `wrk http://localhost:9001`\
Results:

```
{{standard}}
```

## Standard Node.js

Source code: [`node/main.js`](node/main.js)\
Command: `wrk http://localhost:3001`\
Results:

```
{{node}}
```

## Express.js

Source code: [`express/main.js`](express/main.js)\
Command: `wrk http://localhost:3000`\
Results:

```
{{express}}
```
