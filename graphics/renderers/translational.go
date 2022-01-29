package renderers

import (
	"sync"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	ggl "github.com/lucas-s-work/gopengl3/graphics/gl"
	"github.com/lucas-s-work/gopengl3/graphics/gl/shader"
)

const (
	translationVertShader = "./shaders/translational/vertex.vert"
	translationFragShader = "./shaders/translational/frag.frag"
)

const (
	dimensionUniform   = "dimension"
	tranlsationUniform = "trans"
)

type Translational struct {
	*Renderer2D
	shader      *shader.Program
	Translation *mgl32.Vec2
	tMut        sync.Mutex
}

func CreateTranslationalRenderer(window *ggl.Window, texture string, size int32) (*Translational, error) {
	p := shader.CreateProgram(0)
	if err := p.LoadShader(translationVertShader, gl.VERTEX_SHADER); err != nil {
		return nil, err
	}
	if err := p.LoadShader(translationFragShader, gl.FRAGMENT_SHADER); err != nil {
		return nil, err
	}
	if err := p.Link(); err != nil {
		return nil, err
	}

	t := mgl32.Vec2{}
	if err := p.AttachUniform(tranlsationUniform, t); err != nil {
		return nil, err
	}

	r, err := CreateRenderer2D(window, texture, size, p)
	if err != nil {
		return nil, err
	}

	return &Translational{
		Renderer2D:  r,
		shader:      p,
		Translation: &t,
	}, nil
}

func (t *Translational) SetTranslation(translation mgl32.Vec2) {
	t.shader.SetUniform(tranlsationUniform, translation)
}
