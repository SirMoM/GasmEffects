.PHONY: clean serve

GO_FILES := $(shell find . -name "*.go")

buildWasm: clean $(GO_FILES)
	GOOS=js GOARCH=wasm go build -o dist/test.wasm .
	cp wasm_exec.js dist/

serve:
	python3 -m http.server -d "./example"

clean:
	rm -f dist/test.wasm dist/wasm_exec.js
