test:
	golint .
	go test -v .

bench:
	go test -bench . -benchmem

prof:
	go build .
	./length
	go tool pprof --svg length cpu.pprof > cpu.svg

