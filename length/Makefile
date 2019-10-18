test:
	golint .
	go test -v .

bench:
	go test -bench . -benchmem -o test.bin -cpuprofile cpu.pprof
	go tool pprof -svg length cpu.pprof > bench.svg

prof:
	go build .
	./length
	go tool pprof --svg length cpu.pprof > cpu.svg

