package main

import (
	_ "embed"
	"strings"
	"syscall/js"

	"github.com/SirMoM/go-wasm/gasm"
	"github.com/SirMoM/go-wasm/shared"
)

//go:embed VERSION
var version string

func main() {
	js.Global().Set("goAdd", js.FuncOf(add))
	shared.Info("Registered goAdd to JS!")
	js.Global().Set("goMI", js.FuncOf(gasm.ManipulateImg))
	shared.Info("Registered goMI to JS!")
	js.Global().Set("getVersion", js.FuncOf(getVersion))
	shared.Info("Registered getVersion to JS!")
	select {}
}

func getVersion(this js.Value, args []js.Value) any {
	return js.ValueOf(strings.TrimSpace(version))
}

func add(this js.Value, args []js.Value) any {
	var a, b int
	shared.Info(this)
	shared.Info(args)
	a = args[0].Int()
	b = args[1].Int()

	return js.ValueOf(a + b)
}
