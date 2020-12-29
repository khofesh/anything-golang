package main

import (
	"fmt"
	"syscall/js"
)

func add(this js.Value, inputs []js.Value) interface {} {
	js.Global().Set("output", js.ValueOf(inputs[0].Int() + inputs[1].Int()))
	fmt.Println(js.ValueOf(inputs[0].Int() + inputs[1].Int()))

	return nil
}

func subtract(this js.Value, inputs []js.Value) interface {} {
	js.Global().Set("output", js.ValueOf(inputs[0].Int() - inputs[1].Int()))
	fmt.Println(js.ValueOf(inputs[0].Int() - inputs[1].Int()))

	return nil
}

func registerCallbacks() {
	js.Global().Set("add", js.FuncOf(add))
	js.Global().Set("subtract", js.FuncOf(subtract))
}

func main() {
	channel := make(chan struct{}, 0)

	fmt.Println("hello world");

	registerCallbacks();

	<-channel
}