//go:build js && wasm

package main

import (
	"runtime/debug"
	"syscall/js"
)

func main() {
	debug.SetMaxStack(256 << 20) // deep recursion fails fast instead of eating the tab's memory
	js.Global().Set("goJudge", js.FuncOf(func(_ js.Value, args []js.Value) any {
		return judge(args[0].String(), args[1].String(), args[2].String())
	}))
	js.Global().Call("goJudgeReady")
	select {} // keep the program alive for later calls
}
