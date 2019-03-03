all:
	go build 
test:
	go test -v .
bench:
	go test -bench . -benchmem
cover:
	go test -v . -coverprofile=cover.out && go tool cover -html=cover.out -o cover.html
