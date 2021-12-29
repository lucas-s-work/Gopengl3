package vao

import (
	"fmt"

	"github.com/go-gl/gl/v4.1-core/gl"
	ggl "github.com/lucas-s-work/gopengl3/graphics/gl"
	"github.com/lucas-s-work/gopengl3/graphics/gl/shader"
)

type VAO struct {
	id      uint32
	buffers map[string]*Buffer
	window  *ggl.Window
	shader  *shader.Program
	texture *ggl.Texture
	vertNum int32
}

func CreateVAO(window *ggl.Window, textureFile string, shader *shader.Program) (*VAO, error) {
	texture := ggl.LoadTexture(textureFile)

	id, err := ggl.GetFreeVAOIId()
	if err != nil {
		return nil, err
	}

	return &VAO{
		id:      id,
		window:  window,
		texture: texture,
		shader:  shader,
		buffers: make(map[string]*Buffer),
	}, nil
}

func (vao *VAO) Bind() {
	gl.BindVertexArray(vao.id)
}

func (vao *VAO) AttachBuffer(attribute string, size int32) error {
	if _, ok := vao.buffers[attribute]; ok {
		return fmt.Errorf("Unable to attach buffer, already attached with attribute: %s", attribute)
	}

	if n := vao.vertNum; n > 0 && n != size {
		return fmt.Errorf("Unable to attach buffer, size mismatch expected: %v", n)
	}

	vao.Bind()

	b, err := CreateBuffer(attribute, int(size), vao.shader)
	if err != nil {
		return err
	}

	vao.vertNum = size
	vao.buffers[attribute] = b

	return nil
}

func (vao *VAO) SetBuffer(attribute string, elems []float32) error {
	b, ok := vao.buffers[attribute]
	if !ok {
		return fmt.Errorf("No buffer with name: %s attached", attribute)
	}

	lg := len(elems)
	if la := len(b.elements); la != lg {
		return fmt.Errorf("Element length does not match, expected: %v, got: %v", la, lg)
	}

	b.elements = elems
	b.Update()

	return nil
}

func (vao *VAO) SetBufferIndex(attribute string, elems []float32, index int) error {
	b, ok := vao.buffers[attribute]
	if !ok {
		return fmt.Errorf("No buffer with name: %s attached", attribute)
	}

	if index < 0 {
		return fmt.Errorf("Cannot use negative index")
	}

	if len(elems)+index > len(b.elements) {
		return fmt.Errorf("Index + Length larger than allocated elements")
	}

	for i, e := range elems {
		b.elements[i+index] = e
	}

	return nil
}

func (vao *VAO) Shader() *shader.Program {
	return vao.shader
}

func (vao *VAO) Render() {
	vao.shader.Use()
	vao.Bind()
	vao.texture.Use()
	vao.shader.UpdateUniforms()
	gl.DrawArrays(gl.TRIANGLES, 0, vao.vertNum)
}
