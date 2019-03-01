all:
	go build 
test:
	go test -v .

cover:
	go test -v . -coverprofile=cover.out && go tool cover -html=cover.out -o cover.html
