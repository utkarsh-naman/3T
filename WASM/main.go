package main

import (
	"syscall/js"
)

func main() {
	// Create a channel to prevent the Go program from exiting immediately.
	// WASM modules need to stay "running" to listen for JS calls.
	c := make(chan struct{})

	// Register the function on the global 'window' object (js.Global())
	js.Global().Set("TTTNGinPlay", js.FuncOf(wrapperTTTNGinPlay))

	// Block forever so the WASM module stays active
	<-c
}

func wrapperTTTNGinPlay(_ js.Value, args []js.Value) interface{} {
	if len(args) == 0 {
		return "Error: No state provided"
	}

	// 1. Get the input (State) from the first JavaScript argument
	// We use .Int() to get the number, then cast to your State type (uint32).
	inputVal := args[0].Int()
	inputState := State(inputVal)

	// 2. Call your engine function
	resultMove := TTTNGinPlay(inputState)

	// 3. Return the result back to JavaScript
	return int(resultMove)
}
