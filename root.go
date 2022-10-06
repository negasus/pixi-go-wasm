package pixi_go_wasm

import "syscall/js"

func PIXI() js.Value {
	return js.Global().Get("window").Get("PIXI")
}

type DisplayObject interface {
	Value() js.Value
}

type Container interface {
	AddChild(children ...DisplayObject)
	RemoveChild(children ...DisplayObject)
}
