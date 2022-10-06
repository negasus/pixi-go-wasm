package application

import (
	root "github.com/negasus/pixi-go-wasm"
	"github.com/negasus/pixi-go-wasm/assets"
	"syscall/js"
)

type Application struct {
	v      js.Value
	assets *assets.Assets
}

func New(renderTo string, options map[string]any) *Application {
	app := &Application{
		assets: assets.New(),
	}

	app.v = root.PIXI().Get("Application").New(options)

	js.Global().Get("document").Call("getElementById", renderTo).Call("appendChild", app.v.Get("view"))

	return app
}

func (app *Application) AddChild(children ...root.DisplayObject) {
	for _, ch := range children {
		app.v.Get("stage").Call("addChild", ch.Value())
	}
}

func (app *Application) RemoveChild(children ...root.DisplayObject) {
	for _, ch := range children {
		app.v.Get("stage").Call("removeChild", ch.Value())
	}
}

func (app *Application) Assets() *assets.Assets {
	return app.assets
}
