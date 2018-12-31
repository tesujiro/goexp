all:
	go run main.go tables.go binary.go xft.go

test:
	go test -v

bench:
	go test -bench . -benchmem -timeout 0

profile:
	go test -bench . -benchmem -cpuprofile cpu.out
	go tool pprof --svg trie.test cpu.out > cpu.svg
