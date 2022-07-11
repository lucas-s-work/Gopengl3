package renderers

import (
	"fmt"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/lucas-s-work/gopengl3/graphics"
	ggl "github.com/lucas-s-work/gopengl3/graphics/gl"
	"github.com/lucas-s-work/gopengl3/graphics/gl/shader"
	"github.com/lucas-s-work/gopengl3/util"
)

type Renderer2D struct {
	*graphics.BaseRenderer
	vertIndex      int
	vertAssignment util.CompressedList
}

func CreateRenderer2D(window *ggl.Window, texture string, size int32, shader *shader.Program) (*Renderer2D, error) {
	b, err := graphics.CreateBaseRenderer(window, texture, shader)
	if err != nil {
		return nil, err
	}

	v := b.VAO()
	floatNum := size * 2

	err = v.AttachBuffer("verttexcoord", floatNum)
	if err != nil {
		b.Delete()
		return nil, err
	}
	err = v.AttachBuffer("vert", floatNum)
	if err != nil {
		b.Delete()
		return nil, err
	}

	assignments, err := util.CreateCompressedList(int(floatNum))
	if err != nil {
		b.Delete()
		return nil, err
	}

	return &Renderer2D{
		BaseRenderer:   b,
		vertAssignment: *assignments,
	}, nil
}

func (r Renderer2D) SetAttributeValues(attribute string, vertices []mgl32.Vec2, index int) error {
	// unravel the vertices into float array
	elems := make([]float32, len(vertices)*2)
	for i, v := range vertices {
		elems[2*i] = v.X()
		elems[2*i+1] = v.Y()
	}

	if err := r.VAO().SetBufferIndex(attribute, elems, 2*index); err != nil {
		return err
	}

	return nil
}

func (r Renderer2D) AllocateAndSetVertices(verts []mgl32.Vec2, texs []mgl32.Vec2) (*util.ListNode, error) {
	if len(verts) != len(texs) {
		return nil, fmt.Errorf("Unable to allocate and set, vertex and texture coords lengths don't match")
	}

	allocation, err := r.AllocateVertices(len(verts))
	if err != nil {
		return nil, err
	}

	if err := r.SetVertices(verts, texs, allocation); err != nil {
		allocation.Free()
		return nil, err
	}

	return allocation, nil
}

func (r Renderer2D) AllocateVertices(size int) (*util.ListNode, error) {
	n, err := r.vertAssignment.Allocate(size)
	if err != nil {
		return nil, fmt.Errorf("Unable to allocate: %v vertices :%w", size, err)
	}

	return n, nil
}

func (r Renderer2D) SetVertices(verts []mgl32.Vec2, texs []mgl32.Vec2, node *util.ListNode) error {
	if node == nil {
		return fmt.Errorf("node is nil")
	}

	if len(verts) != len(texs) {
		return fmt.Errorf("Cannot set vertices, mismatched vertex and texture array dimensions")
	}

	if len(verts) > node.Size() {
		return fmt.Errorf("Cannot set more vertices than allocated")
	}

	if err := r.SetAttributeValues("vert", verts, node.Index()); err != nil {
		return err
	}
	return r.SetAttributeValues("verttexcoord", texs, node.Index())
}

func (r Renderer2D) SetSubVertices(verts []mgl32.Vec2, texs []mgl32.Vec2, node *util.ListNode, subIndex int) error {
	if node == nil {
		return fmt.Errorf("node is nil")
	}

	if len(verts) != len(texs) {
		return fmt.Errorf("Cannot set subvertices vertices, mismatched vertex and texture array dimensions")
	}

	if len(verts)+subIndex > node.Size() {
		return fmt.Errorf("Cannot set subvertices, more vertices than allocated when including subIndex")
	}

	if err := r.SetAttributeValues("vert", verts, node.Index()+subIndex); err != nil {
		return err
	}
	return r.SetAttributeValues("verttexcoord", texs, node.Index()+subIndex)
}

func (r Renderer2D) ClearVertices(node *util.ListNode) error {
	if node == nil {
		return fmt.Errorf("node is nil")
	}
	err := r.SetVertices(make([]mgl32.Vec2, node.Size()), make([]mgl32.Vec2, node.Size()), node)
	if err != nil {
		return err
	}

	node.Free()

	return nil
}
