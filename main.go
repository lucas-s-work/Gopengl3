package main

import (
	"runtime"

	"github.com/go-gl/mathgl/mgl32"
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

	r, err := renderers.CreateTranslationalRenderer(window, "./textures/test.png", 12)
	if err != nil {
		panic(err)
	}

	ctx.Attach(r, 1)

	u, err := renderers.CreateTranslationalRenderer(window, "./textures/test.png", 12)

	ctx.Attach(u, 0)

	verts, texs, _ := graphics.Square(0, 0, 32, 0, 0, 1, r.Texture(), window)
	r.AllocateAndSetVertices(verts, texs)
	verts, texs, _ = graphics.Square(0, 0, 32, 1, 0, 1, u.Texture(), window)
	u.AllocateAndSetVertices(verts, texs)
	u.SetTranslation(mgl32.Vec2{0.02, 0})

	r.Update()
	u.Update()

	for !window.ShouldClose() {
		ctx.Render()
	}
}
