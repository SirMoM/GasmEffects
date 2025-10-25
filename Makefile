.PHONY: clean

buildWasm: clean main.go jsparser.go imageManipulation.go jsUtils.go
	GOOS=js GOARCH=wasm go build -o dist/test.wasm
	cp wasm_exec.js dist/

clean:
	rm -f dist/test.wasm dist/wasm_exec.js

