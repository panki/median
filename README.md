# Find median two sorted arrays of same sizes

Run benchmark:

```
make test
```

Mac mini (M1) results:

```
go test -bench=. -benchmem -benchtime 10000x
goos: darwin
goarch: arm64
pkg: github.com/panki/median
BenchmarkMedian_ON-8      	   10000	    362540 ns/op	  401409 B/op	       1 allocs/op
BenchmarkMedian_ON2-8     	   10000	    161380 ns/op	       0 B/op	       0 allocs/op
BenchmarkMedian_OLOGN-8   	   10000	        44.29 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/panki/median	5.344s
```

Macbook Pro 2017 results:

````
go test -bench=. -benchmem -benchtime 10000x
goos: darwin
goarch: amd64
pkg: github.com/panki/median
cpu: Intel(R) Core(TM) i7-7820HQ CPU @ 2.90GHz
BenchmarkMedian_ON-8      	   10000	    585645 ns/op	  401409 B/op	       1 allocs/op
BenchmarkMedian_ON2-8     	   10000	    229739 ns/op	       0 B/op	       0 allocs/op
BenchmarkMedian_OLOGN-8   	   10000	        57.82 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/panki/median	8.373s```
````
