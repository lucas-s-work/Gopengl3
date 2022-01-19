package graphics

import (
	"github.com/lucas-s-work/gopengl3/graphics/gl"
	"github.com/lucas-s-work/gopengl3/graphics/gl/shader"
	"github.com/lucas-s-work/gopengl3/graphics/gl/vao"
)

type Renderer interface {
	VAO() *vao.VAO
	Texture() *gl.Texture
	Update()
	Delete()
	Render()
}

type BaseRenderer struct {
	shader *shader.Program
	vao    *vao.VAO
}

func CreateBaseRenderer(window *gl.Window, texture string, shader *shader.Program) (*BaseRenderer, error) {
	vao, err := vao.CreateVAO(window, texture, shader)
	if err != nil {
		return nil, err
	}

	r := &BaseRenderer{
		shader: shader,
		vao:    vao,
	}

	return r, nil
}

func (r BaseRenderer) VAO() *vao.VAO {
	return r.vao
}

func (r BaseRenderer) Update() {
	r.VAO().UpdateBuffers()
}

func (r BaseRenderer) Delete() {
	r.vao.Delete()
	r.shader.Delete()
}

func (r BaseRenderer) Render() {
	r.vao.Render()
}

func (r BaseRenderer) Texture() *gl.Texture {
	return r.VAO().Texture()
}