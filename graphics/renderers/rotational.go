package renderers

import (
	"sync"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	ggl "github.com/lucas-s-work/gopengl3/graphics/gl"
	"github.com/lucas-s-work/gopengl3/graphics/gl/shader"
)

const (
	rotationVertShader = "./shaders/rotational/vertex.vert"
	rotationFragShader = "./shaders/rotational/frag.frag"
)

const (
	rotation1AngleUniform  = "rot1angle"
	rotation1CenterUniform = "rot1center"
)

type Rotational struct {
	*Renderer2D
	shader *shader.Program
	tMut   sync.Mutex
}

func CreateRotationalRenderer(window *ggl.Window, texture string, size int32) (*Rotational, error) {
	p := shader.CreateProgram(0)
	if err := p.LoadShader(rotationVertShader, gl.VERTEX_SHADER); err != nil {
		return nil, err
	}
	if err := p.LoadShader(rotationFragShader, gl.FRAGMENT_SHADER); err != nil {
		return nil, err
	}
	if err := p.Link(); err != nil {
		return nil, err
	}

	t := mgl32.Vec2{}
	if err := p.AttachUniform(tranlsationUniform, t); err != nil {
		return nil, err
	}
	var rAngle float32 = 0
	if err := p.AttachUniform(rotation1AngleUniform, rAngle); err != nil {
		return nil, err
	}
	rCenter := mgl32.Vec2{}
	if err := p.AttachUniform(rotation1CenterUniform, rCenter); err != nil {
		return nil, err
	}
	if err := p.AttachUniform(dimensionUniform, mgl32.Vec2{window.Width, window.Height}); err != nil {
		return nil, err
	}

	r, err := CreateRenderer2D(window, texture, size, p)
	if err != nil {
		return nil, err
	}

	return &Rotational{
		Renderer2D: r,
		shader:     p,
	}, nil
}

func (r *Rotational) SetTranslation(translation mgl32.Vec2) {
	r.shader.SetUniform(tranlsationUniform, translation)
}

func (r *Rotational) SetRotation1(angle float32, center mgl32.Vec2) {
	r.shader.SetUniform(rotation1AngleUniform, angle)
	r.shader.SetUniform(rotation1CenterUniform, center)
}
