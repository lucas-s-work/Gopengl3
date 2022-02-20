package graphics

import (
	"sync"

	"github.com/lucas-s-work/gopengl3/graphics/gl"
	"github.com/lucas-s-work/gopengl3/graphics/gl/shader"
	"github.com/lucas-s-work/gopengl3/graphics/gl/vao"
)

type Renderer interface {
	VAO() *vao.VAO
	Texture() *gl.Texture
	SetTexture(*gl.Texture)
	Update()
	SetId(int)
	GetId() int
	Delete()
	SetLayer(int)
	GetLayer() int
	SetActive(bool)
	GetActive() bool
	Render()
}

type BaseRenderer struct {
	shader    *shader.Program
	vao       *vao.VAO
	id, layer int
	active    bool
	mut       sync.Mutex
}

func CreateBaseRenderer(window *gl.Window, texture string, shader *shader.Program) (*BaseRenderer, error) {
	vao, err := vao.CreateVAO(window, texture, shader)
	if err != nil {
		return nil, err
	}

	r := &BaseRenderer{
		shader: shader,
		vao:    vao,
		active: true,
	}

	return r, nil
}

func (r *BaseRenderer) VAO() *vao.VAO {
	return r.vao
}

func (r *BaseRenderer) SetId(id int) {
	r.id = id
}

func (r *BaseRenderer) GetId() int {
	return r.id
}

func (r *BaseRenderer) SetLayer(layer int) {
	r.layer = layer
}
func (r *BaseRenderer) GetLayer() int {
	return r.layer
}

func (r *BaseRenderer) Update() {
	r.VAO().UpdateBuffers()
}

func (r *BaseRenderer) Delete() {
	r.vao.Delete()
	r.shader.Delete()
}

func (r *BaseRenderer) Render() {
	// Use the function call here to simplify mutex
	if !r.GetActive() {
		return
	}
	r.vao.Render()
}

func (r *BaseRenderer) Texture() *gl.Texture {
	return r.VAO().Texture()
}

func (r *BaseRenderer) SetTexture(t *gl.Texture) {
	r.VAO().SetTexture(t)
}

func (r *BaseRenderer) SetActive(a bool) {
	r.mut.Lock()
	defer r.mut.Unlock()
	r.active = a
}

func (r *BaseRenderer) GetActive() bool {
	r.mut.Lock()
	defer r.mut.Unlock()
	return r.active
}
