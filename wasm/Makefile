build:
	GOOS=js GOARCH=wasm go build -o ./public/main.wasm

cp_js:
	cp /usr/local/go/misc/wasm/wasm_exec.js ./public

start_server:
	./server/server &

