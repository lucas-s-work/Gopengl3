package renderers

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/lucas-s-work/gopengl3/graphics"
	ggl "github.com/lucas-s-work/gopengl3/graphics/gl"
	"github.com/lucas-s-work/gopengl3/graphics/gl/shader"
)

const (
	vertShader = "./shaders/translational/vertex.vert"
	fragShader = "./shaders/translational/frag.frag"
)

type Translational struct {
	*graphics.BaseRenderer
	shader      *shader.Program
	Translation *mgl32.Vec2
}

func CreateTranslationalRenderer(window *ggl.Window, texture string, size int32) (*Translational, error) {
	p := shader.CreateProgram(0)
	if err := p.LoadShader(vertShader, gl.VERTEX_SHADER); err != nil {
		return nil, err
	}
	if err := p.LoadShader(fragShader, gl.FRAGMENT_SHADER); err != nil {
		return nil, err
	}
	if err := p.Link(); err != nil {
		return nil, err
	}

	t := mgl32.Vec2{}
	p.AttachUniform("trans", t)

	b, err := graphics.CreateBaseRenderer(window, texture, p)
	if err != nil {
		return nil, err
	}

	v := b.VAO()
	v.AttachBuffer("vert", size)
	v.AttachBuffer("verttexcoord", size)

	return &Translational{
		BaseRenderer: b,
		shader:       p,
		Translation:  &t,
	}, nil
}
