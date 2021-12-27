package shader

import (
	"fmt"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type Uniform struct {
	id      int32
	value   interface{}
	updated bool
}

func (u *Uniform) Update() error {
	if !u.updated {
		return nil
	}

	if u.id == -1 {
		return fmt.Errorf("Uniform not found")
	}

	switch t := u.value.(type) {
	case float32:
		gl.Uniform1f(u.id, (u.value).(float32))
	case mgl32.Vec2:
		v := (u.value).(mgl32.Vec2)
		gl.Uniform2f(u.id, v.X(), v.Y())
	case mgl32.Vec3:
		v := (u.value).(mgl32.Vec3)
		gl.Uniform3f(u.id, v.X(), v.Y(), v.Z())
	case mgl32.Mat2:
		v := (u.value).(mgl32.Mat2)
		gl.UniformMatrix2fv(u.id, 1, false, &v[0])
	case mgl32.Mat3:
		v := (u.value).(mgl32.Mat3)
		gl.UniformMatrix3fv(u.id, 1, false, &v[0])
	default:
		u.updated = false
		return fmt.Errorf("Unsupported uniform type :%v", t)
	}

	u.updated = false
	return nil
}

func (u *Uniform) Set(value interface{}) {
	u.updated = true
	u.value = value
}
