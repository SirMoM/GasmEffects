buildWams: main.go jsparser.go imageManipulation.go jsUtils.go
	GOOS=js GOARCH=wasm go build -o test.wasm
