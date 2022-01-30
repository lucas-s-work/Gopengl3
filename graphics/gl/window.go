package gl

import (
	"github.com/go-gl/glfw/v3.2/glfw"
)

type Window struct {
	glWindow      *glfw.Window
	Width, Height float32
	Name          string
}

type InputInfo struct {
	Mx, My     int
	M1, M2, M3 bool
	KeyMap     map[string]bool
}

func CreateWindow(width, height int, name string) (*Window, error) {
	if err := glfw.Init(); err != nil {
		return nil, err
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(width, height, name, nil, nil)
	if err != nil {
		return nil, err
	}

	window.MakeContextCurrent()

	return &Window{
		glWindow: window,
		Width:    float32(width),
		Height:   float32(height),
		Name:     name,
	}, nil
}

func (w *Window) Destroy() {
	w.glWindow.Destroy()
}

func (w *Window) SwapBuffers() {
	w.glWindow.SwapBuffers()
}

func (w *Window) PollInput() {
	glfw.PollEvents()
	updateKeys(w.glWindow)
}

func (w *Window) ShouldClose() bool {
	return w.glWindow.ShouldClose()
}

func (w *Window) ScreenToPix(x, y float32) (int, int) {
	// Bottom left is 0,0
	x++
	y++

	x *= w.Width / 2
	y *= w.Height / 2

	return int(x), int(y)
}

func (w *Window) PixToScreen(x, y int) (float32, float32) {
	sx := (2 * float32(x)) / w.Width
	sy := (2 * float32(y)) / w.Height

	sx--
	sy--

	return sx, sy
}
