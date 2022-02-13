package gl

import (
	"sync"

	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

type Window struct {
	glWindow      *glfw.Window
	Width, Height float32
	Name          string
}

type InputInfo struct {
	Mx, My     int
	M1, M2, M3 key
	KeyMap     map[string]key
}

type key struct {
	KeyState, LastKeyState bool
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
	w.updateKeys()
	w.updateMouse()
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

var (
	keyMap   = map[glfw.Key]*key{}
	mouseMap = map[glfw.MouseButton]*key{}
	keyMut   = sync.Mutex{}
	mousePos = mgl32.Vec2{}

	keys = []glfw.Key{
		glfw.KeyW,
		glfw.KeyD,
		glfw.KeyS,
		glfw.KeyA,
		glfw.KeyUp,
		glfw.KeyDown,
		glfw.KeyRight,
		glfw.KeyLeft,
	}

	mouseButtons = []glfw.MouseButton{
		glfw.MouseButton1,
		glfw.MouseButton2,
		glfw.MouseButton3,
	}
)

func (w *Window) updateKeys() {
	keyMut.Lock()
	defer keyMut.Unlock()

	for _, k := range keys {
		w.checkKeyState(k)
	}
}

func (w *Window) checkKeyState(glfwKey glfw.Key) {
	if k, ok := keyMap[glfwKey]; ok {
		k.KeyState = w.glWindow.GetKey(glfwKey) == glfw.Press
	} else {
		keyMap[glfwKey] = &key{
			KeyState: w.glWindow.GetKey(glfwKey) == glfw.Press,
		}
	}
}

func (w *Window) checkMouseState(glfwButton glfw.MouseButton) {
	if k, ok := mouseMap[glfwButton]; ok {
		k.KeyState = w.glWindow.GetMouseButton(glfwButton) == glfw.Press
	} else {
		mouseMap[glfwButton] = &key{
			KeyState: w.glWindow.GetMouseButton(glfwButton) == glfw.Press,
		}
	}
}

func (w *Window) updateMouse() {
	keyMut.Lock()
	defer keyMut.Unlock()

	window := w.glWindow

	x, y := window.GetCursorPos()
	mousePos = mgl32.Vec2{float32(x), w.Height - float32(y)}

	for _, b := range mouseButtons {
		w.checkMouseState(b)
	}
}

func GetMouseInfo() (mgl32.Vec2, map[glfw.MouseButton]*key) {
	keyMut.Lock()
	defer keyMut.Unlock()

	return mousePos, mouseMap
}

func CheckKeyPressed(key glfw.Key) bool {
	keyMut.Lock()
	defer keyMut.Unlock()

	if keyMap[key] == nil {
		return false
	}

	return keyMap[key].KeyState
}

func CheckKeyTapped(key glfw.Key) bool {
	keyMut.Lock()
	defer keyMut.Unlock()

	s := keyMap[key]
	if s == nil {
		return false
	}

	tapped := s.KeyState && (s.KeyState != s.LastKeyState)
	s.LastKeyState = s.KeyState
	return tapped
}

func CheckMouseTapped(button glfw.MouseButton) bool {
	keyMut.Lock()
	defer keyMut.Unlock()

	s := mouseMap[button]
	if s == nil {
		return false
	}

	tapped := s.KeyState && (s.KeyState != s.LastKeyState)
	s.LastKeyState = s.KeyState
	return tapped
}

func CheckKeysPressed(keys []glfw.Key) []bool {
	keyMut.Lock()
	defer keyMut.Unlock()

	out := make([]bool, len(keys))
	for i, k := range keys {
		out[i] = keyMap[k].KeyState
	}

	return out
}

func CheckKeyComboPressed(keys []glfw.Key) bool {
	keyMut.Lock()
	defer keyMut.Unlock()

	for _, v := range keys {
		k, ok := keyMap[v]
		if !k.KeyState || !ok {
			return false
		}
	}

	return true
}
