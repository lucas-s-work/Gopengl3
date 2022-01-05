package renderers

import (
	"fmt"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/lucas-s-work/gopengl3/graphics"
	ggl "github.com/lucas-s-work/gopengl3/graphics/gl"
	"github.com/lucas-s-work/gopengl3/graphics/gl/shader"
)

type Renderer2D struct {
	*graphics.BaseRenderer
	vertIndex int
}

func CreateRenderer2D(window *ggl.Window, texture string, size int32, shader *shader.Program) (*Renderer2D, error) {
	b, err := graphics.CreateBaseRenderer(window, texture, shader)
	if err != nil {
		return nil, err
	}

	v := b.VAO()
	v.AttachBuffer("vert", size)
	v.AttachBuffer("verttexcoord", size)

	return &Renderer2D{
		BaseRenderer: b,
	}, nil
}

func (r Renderer2D) SetAttributeValues(attribute string, vertices []mgl32.Vec2, index int) error {
	// unravel the vertices into float array
	elems := make([]float32, len(vertices)*2)
	for i, v := range vertices {
		elems[2*i] = v.X()
		elems[2*i+1] = v.Y()
	}

	if err := r.VAO().SetBufferIndex(attribute, elems, index); err != nil {
		return err
	}

	return nil
}

func (r Renderer2D) SetVertices(verts []mgl32.Vec2, texs []mgl32.Vec2, index int) error {
	if len(verts) != len(texs) {
		return fmt.Errorf("Cannot set vertices, mismatched vertex and texture array dimensions")
	}

	if err := r.SetAttributeValues("vert", verts, index); err != nil {
		return err
	}
	return r.SetAttributeValues("verttexcoord", texs, index)
}

func (r Renderer2D) ClearVertices(index, size int) error {
	if index < 0 || size < 0 {
		return fmt.Errorf("invalid arguements given, < 0")
	}

	// TODO use a linked list to compress which set of coordinates are currently in use and which arent

	return r.SetVertices(make([]mgl32.Vec2, size), make([]mgl32.Vec2, size), index)
}
