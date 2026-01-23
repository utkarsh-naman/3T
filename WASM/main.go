package main

import "syscall/js"

func tttNginPlayWrapper(_ js.Value, args []js.Value) any {
	if len(args) != 1 {
		return js.ValueOf("Expected 1 argument: state (uint32)")
	}

	state := State(uint32(args[0].Int()))
	move := TTTNGinPlay(state)

	return js.ValueOf(int(move))
}

func main() {
	js.Global().Set("TTTNGinPlay", js.FuncOf(tttNginPlayWrapper))
	select {} // keep WASM alive
}
