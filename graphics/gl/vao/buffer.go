package vao

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/lucas-s-work/gopengl3/graphics/gl/shader"
)

type Buffer struct {
	id        uint32
	elements  []float32
	attribute string
	size      int
}

func CreateBuffer(attribute string, size int, shader *shader.Program) (*Buffer, error) {
	buffer := &Buffer{
		elements:  make([]float32, size),
		attribute: attribute,
		size:      size,
	}

	gl.GenBuffers(1, &buffer.id)
	gl.BindBuffer(gl.ARRAY_BUFFER, buffer.id)
	gl.BufferData(gl.ARRAY_BUFFER, 2*size, gl.Ptr(buffer.elements), gl.DYNAMIC_DRAW)

	if err := shader.AttachAttribute(attribute, 2); err != nil {
		buffer.Delete()

		return nil, err
	}

	return buffer, nil
}

func (buffer *Buffer) Update() {
	gl.BindBuffer(gl.ARRAY_BUFFER, buffer.id)
	gl.BufferSubData(gl.ARRAY_BUFFER, 0, 2*buffer.size, gl.Ptr(buffer.elements))
}

func (buffer *Buffer) Delete() {
	gl.DeleteBuffers(1, &buffer.id)
}

func (buffer *Buffer) VertNum() int {
	return buffer.size
}
