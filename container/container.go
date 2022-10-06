package container

import "syscall/js"

type Container struct {
	v js.Value
}

func New(v js.Value) *Container {
	return &Container{
		v: v,
	}
}
