package sprite

import (
	"syscall/js"
)

type Sprite struct {
	v js.Value
}

func From(texture string) *Sprite {
	s := &Sprite{}

	s.v = js.Global().Get("window").Get("PIXI").Get("Sprite").Call("from", texture)

	return s
}

func (s *Sprite) Value() js.Value {
	return s.v
}

func (s *Sprite) SetX(x int) *Sprite {
	s.v.Set("x", x)
	return s
}

func (s *Sprite) SetY(y int) *Sprite {
	s.v.Set("y", y)
	return s
}
