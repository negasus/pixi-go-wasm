package text

import (
	"syscall/js"
)

type Text struct {
	v js.Value
}

func New(value string) *Text {
	p := &Text{}

	p.v = js.Global().Get("window").Get("PIXI").Get("Text").New(value)

	return p
}

func (t *Text) Value() js.Value {
	return t.v
}

func (t *Text) SetX(x int) *Text {
	t.v.Set("x", x)
	return t
}

func (t *Text) SetY(y int) *Text {
	t.v.Set("y", y)
	return t
}
