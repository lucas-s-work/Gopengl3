package gl

import (
	"sync"

	"github.com/go-gl/glfw/v3.2/glfw"
)

var (
	keyMap = map[string]bool{}
	keyMut = sync.Mutex{}
)

func updateKeys(window *glfw.Window) {
	keyMut.Lock()
	defer keyMut.Unlock()

	keyMap["w"] = window.GetKey(glfw.KeyW) == glfw.Press
	keyMap["a"] = window.GetKey(glfw.KeyA) == glfw.Press
	keyMap["s"] = window.GetKey(glfw.KeyS) == glfw.Press
	keyMap["d"] = window.GetKey(glfw.KeyD) == glfw.Press

	keyMap["up"] = window.GetKey(glfw.KeyUp) == glfw.Press
	keyMap["down"] = window.GetKey(glfw.KeyDown) == glfw.Press
	keyMap["left"] = window.GetKey(glfw.KeyLeft) == glfw.Press
	keyMap["right"] = window.GetKey(glfw.KeyRight) == glfw.Press
}

func CheckKey(key string) bool {
	keyMut.Lock()
	defer keyMut.Unlock()

	return keyMap[key]
}

func CheckKeys(keys []string) []bool {
	keyMut.Lock()
	defer keyMut.Unlock()

	out := make([]bool, len(keys))
	for i, k := range keys {
		out[i] = keyMap[k]
	}

	return out
}

func CheckKeyCombo(keys []string) bool {
	keyMut.Lock()
	defer keyMut.Unlock()

	for _, v := range keys {
		k, ok := keyMap[v]
		if !k || !ok {
			return false
		}
	}

	return true
}
