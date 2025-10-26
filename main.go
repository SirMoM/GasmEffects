package main

import (
	gasm "github.com/SirMoM/go-wasm/gasm"
	"syscall/js"
)

func main() {
	js.Global().Set("goAdd", js.FuncOf(add))
	gasm.Info("Registered goAdd to JS!")
	js.Global().Set("goMI", js.FuncOf(gasm.ManipulateImg))
	gasm.Info("Registered goMI to JS!")
	select {}
}

func add(this js.Value, args []js.Value) any {
	var a, b int
	gasm.Info(this)
	gasm.Info(args)
	a = args[0].Int()
	b = args[1].Int()

	return js.ValueOf(a + b)
}
