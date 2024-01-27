run:
	go run .

wasm-run:
	GOOS=js GOARCH=wasm go run . -exec="$(shell go env GOROOT)/misc/wasm/go_js_wasm_exec"

wasm-build:
	GOOS=js GOARCH=wasm go build -o main.wasm

wasm-build-tiny:
	GOOS=js GOARCH=wasm tinygo build -o main.wasm

.PHONY: wasm-build wasm-build-tiny wasm-run run
