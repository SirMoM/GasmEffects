package main

import (
	"syscall/js"
)

func main() {
	js.Global().Set("goAdd", js.FuncOf(add))
	info("Registered goAdd to JS!")
	js.Global().Set("goMI", js.FuncOf(manipulateImg))
	info("Registered goMI to JS!")
	select {}
}

func add(this js.Value, args []js.Value) any {
	var a, b int
	info(this)
	info(args)
	a = args[0].Int()
	b = args[1].Int()

	return js.ValueOf(a + b)
}
