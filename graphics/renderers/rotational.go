package renderers

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/lucas-s-work/gopengl3/graphics/gl/shader"
	"github.com/lucas-s-work/gopengl3/graphics/gl/vao"
)

type Rotational struct {
}

func CreateRotationRenderer() {
	p := shader.CreateProgram(0)
	p.LoadShader("./shader.vert", gl.VERTEX_SHADER)
}

func (r Rotational) VAO() *vao.VAO {

	return nil
}

func (r Rotational) Delete() {

}

func (r Rotational) Render() {

}
