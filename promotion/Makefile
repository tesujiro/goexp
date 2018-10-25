.PHONY: test
test:
	go test -v .

.PHONY: bench
bench:
	go test -bench . -benchmem

.PHONY: profile
profile:
	go test -bench . -benchmem -cpuprofile cpu.out
	go tool pprof --svg promotion.test cpu.out > cpu.svg
