package main

import (
	"runtime"

	"github.com/lucas-s-work/gopengl3/graphics/gl"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	window, err := gl.CreateWindow(800, 600, "test")
	if err != nil {
		panic(err)
	}

	if err := gl.GlInit(); err != nil {
		panic(err)
	}
}
