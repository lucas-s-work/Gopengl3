package main

import (
	"runtime"

	"github.com/lucas-s-work/gopengl3/graphics"
	"github.com/lucas-s-work/gopengl3/graphics/gl"
	"github.com/lucas-s-work/gopengl3/graphics/renderers"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	setupOpengl()
}

func setupOpengl() {
	window, err := gl.CreateWindow(800, 600, "test")
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	if err := gl.GlInit(); err != nil {
		panic(err)
	}

	ctx := graphics.CreateContext(window)
	defer ctx.Delete()

	r, err := renderers.CreateTranslationalRenderer(window, "./textures/test.png", 4)
	if err != nil {
		panic(err)
	}

	ctx.Attach(r, 0)

	for !window.ShouldClose() {
		ctx.Render()
	}
}
