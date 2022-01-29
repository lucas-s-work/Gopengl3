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

	var r *renderers.Translational
	ctx.AddJob(func() {
		r, err = renderers.CreateTranslationalRenderer(window, "./textures/test.png", 12)
		if err != nil {
			panic(err)
		}

		v, t, err := graphics.Square(0, 0, 32, 0, 0, 1, r.Texture())
		if err != nil {
			panic(err)
		}

		r.AllocateAndSetVertices(v, t)
		r.SetTranslation(mgl32.Vec2{300, 300})
		// r.SetRotation1(0.5, mgl32.Vec2{300, 300})
		r.Update()
		ctx.Attach(r, 0)
	})

	doneChan := start(window, ctx)

	for !window.ShouldClose() {
		ctx.Render()
	}

	doneChan <- struct{}{}
}

func start(window *gl.Window, ctx *graphics.Context) chan<- struct{} {
	doneChan := make(chan struct{})
	ctxSync := ctx.GetSync()
	go func() {
		for {
			select {
			case <-doneChan:
				return
			case ctxSync <- struct{}{}:
			}
		}
	}()

	return doneChan
}
