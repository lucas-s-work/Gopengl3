package renderers

import "github.com/lucas-s-work/gopengl3/graphics/gl/vao"

type Translational struct {
}

func (r Translational) VAO() *vao.VAO {
	return nil
}

func (r Translational) Delete() {

}

func (r Translational) Render() {

}
