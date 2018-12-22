# What is the purpose of the benchmark
Measure a perfomance of checking existence of a element in the list.

# Compare what?

* Binary Search
* Hashtable (map)

# Environment
```
MacBook Pro (13-inch, 2016, Four Thunderbolt 3 Ports)
Processor 2.9 GHz Intel Core i5
Memory 16 GB 2133 MHz LPDDR3
macOS 10.14.1

go version go1.11 darwin/amd64
```

# Benchmark Results
```
go test -bench . -benchmem -timeout 0
goos: darwin
goarch: amd64
pkg: github.com/tesujiro/goexp/benchFindFromSet
BenchmarkBinary/FindFrom2^0-4         	100000000	        10.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkBinary/FindFrom2^1-4         	100000000	        15.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkBinary/FindFrom2^2-4         	100000000	        19.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkBinary/FindFrom2^3-4         	100000000	        24.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkBinary/FindFrom2^4-4         	50000000	        28.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkBinary/FindFrom2^5-4         	50000000	        34.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkBinary/FindFrom2^6-4         	50000000	        40.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkBinary/FindFrom2^7-4         	30000000	        47.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkBinary/FindFrom2^8-4         	30000000	        56.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkBinary/FindFrom2^9-4         	20000000	        62.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkBinary/FindFrom2^10-4        	20000000	        67.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkBinary/FindFrom2^11-4        	20000000	        74.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkBinary/FindFrom2^12-4        	20000000	        78.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkBinary/FindFrom2^13-4        	20000000	        87.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkBinary/FindFrom2^14-4        	20000000	        96.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkBinary/FindFrom2^15-4        	20000000	       110 ns/op	       0 B/op	       0 allocs/op
BenchmarkBinary/FindFrom2^16-4        	10000000	       135 ns/op	       0 B/op	       0 allocs/op
BenchmarkBinary/FindFrom2^17-4        	10000000	       164 ns/op	       0 B/op	       0 allocs/op
BenchmarkBinary/FindFrom2^18-4        	10000000	       184 ns/op	       0 B/op	       0 allocs/op
BenchmarkBinary/FindFrom2^19-4        	 5000000	       226 ns/op	       0 B/op	       0 allocs/op
BenchmarkHashtable/FindFrom2^0-4      	100000000	        10.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkHashtable/FindFrom2^1-4      	100000000	        10.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkHashtable/FindFrom2^2-4      	100000000	        10.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkHashtable/FindFrom2^3-4      	100000000	        15.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkHashtable/FindFrom2^4-4      	100000000	        15.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkHashtable/FindFrom2^5-4      	100000000	        15.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkHashtable/FindFrom2^6-4      	100000000	        15.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkHashtable/FindFrom2^7-4      	100000000	        15.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkHashtable/FindFrom2^8-4      	100000000	        16.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkHashtable/FindFrom2^9-4      	100000000	        16.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkHashtable/FindFrom2^10-4     	100000000	        16.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkHashtable/FindFrom2^11-4     	100000000	        17.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkHashtable/FindFrom2^12-4     	100000000	        18.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkHashtable/FindFrom2^13-4     	50000000	        20.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkHashtable/FindFrom2^14-4     	50000000	        21.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkHashtable/FindFrom2^15-4     	50000000	        24.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkHashtable/FindFrom2^16-4     	50000000	        25.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkHashtable/FindFrom2^17-4     	50000000	        29.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkHashtable/FindFrom2^18-4     	50000000	        41.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkHashtable/FindFrom2^19-4     	30000000	        51.6 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/tesujiro/goexp/benchFindFromSet	141.301s
```

# Conclusion
Use map if length of list is more than 2^2

