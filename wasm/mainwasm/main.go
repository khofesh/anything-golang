package main

import (
	"fmt"
	"syscall/js"
	"strconv"
)

func addOrSubtract(opType string, inputs []js.Value) (interface {}, error) {
	value1 := js.Global().Get("document").Call("getElementById", inputs[0]).Get("value").String()
	value2 := js.Global().Get("document").Call("getElementById", inputs[1]).Get("value").String()


	int1, error := strconv.Atoi(value1)
	if error != nil {
		return nil, error
	}

	int2, error := strconv.Atoi(value2)
	if error != nil {
		return nil, error
	}

	if opType == "add" {
		return int1 + int2, nil
	} else if opType == "subtract" {
		return int1 - int2, nil
	}

	return nil, nil
}

func add(this js.Value, inputs []js.Value) interface {} {
	result, error := addOrSubtract("add", inputs)
	if error != nil {
		fmt.Println(error)
	}

	// js.Global().Set("output", js.ValueOf(inputs[0].Int() + inputs[1].Int()))
	js.Global().Get("document").Call("getElementById", inputs[2].String()).Set("value", result)

	fmt.Println(result)

	return nil
}

func subtract(this js.Value, inputs []js.Value) interface {} {
	result, error := addOrSubtract("subtract", inputs)
	if error != nil {
		fmt.Println(error)
	}

	// js.Global().Set("output", js.ValueOf(inputs[0].Int() + inputs[1].Int()))
	js.Global().Get("document").Call("getElementById", inputs[2].String()).Set("value", result)

	fmt.Println(result)

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