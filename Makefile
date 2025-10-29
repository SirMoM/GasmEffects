.PHONY: clean serve

GO_FILES := $(shell find . -name "*.go")
VERSION_FILE := VERSION
VERSION := $(shell git rev-parse --short HEAD)
DIRTY := $(shell git status --porcelain)

version:
	@echo "$(VERSION)$(if $(DIRTY),-dirty)" > $(VERSION_FILE)

buildWasm: clean version $(GO_FILES)
	GOOS=js GOARCH=wasm go build -o dist/test.wasm .
	cp wasm_exec.js dist/
	cp $(VERSION_FILE) dist/
	cp -r dist example/dist

serve:
	python3 -m http.server -d "./example"

clean:
	rm -f dist/test.wasm dist/wasm_exec.js dist/VERSION
	rm -rf example/dist
	rm -f $(VERSION_FILE)
