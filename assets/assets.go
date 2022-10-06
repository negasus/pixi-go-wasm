package assets

import (
	root "github.com/negasus/pixi-go-wasm"
	"syscall/js"
)

type Assets struct {
	sharedLoader js.Value
}

func New() *Assets {
	a := &Assets{
		sharedLoader: root.PIXI().Get("Loader").Get("shared"),
	}

	return a
}

func (a *Assets) Add(name, path string) *Assets {
	a.sharedLoader.Call("add", name, path)
	return a
}

func (a *Assets) Load(callback func(args []js.Value)) {
	a.sharedLoader.Call("load", js.FuncOf(func(this js.Value, args []js.Value) any {
		callback(args)
		return nil
	}))
}
