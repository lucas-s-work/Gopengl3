package main

import (
	"runtime"

	"github.com/lucas-s-work/gopengl3/graphics/gl"
	"github.com/lucas-s-work/gopengl3/graphics/renderers"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	_, err := gl.CreateWindow(800, 600, "test")
	if err != nil {
		panic(err)
	}

	if err := gl.GlInit(); err != nil {
		panic(err)
	}

	renderers.CreateRotationRenderer()
}
