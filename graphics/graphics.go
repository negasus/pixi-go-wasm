package graphics

import "syscall/js"

type Graphics struct {
	v js.Value
}

func New() *Graphics {
	g := &Graphics{}

	g.v = js.Global().Get("window").Get("PIXI").Get("Graphics").New()

	return g
}

func (g *Graphics) Value() js.Value {
	return g.v
}

func (g *Graphics) BeginFill(color int) *Graphics {
	g.v.Call("beginFill", color)
	return g
}

func (g *Graphics) DrawRect(x, y, width, height int) *Graphics {
	g.v.Call("drawRect", x, y, width, height)
	return g
}

func (g *Graphics) EndFill() *Graphics {
	g.v.Call("endFill")
	return g
}
