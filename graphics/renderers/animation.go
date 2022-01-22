package renderers

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	ggl "github.com/lucas-s-work/gopengl3/graphics/gl"
	"github.com/lucas-s-work/gopengl3/util"
)

type Animation2D struct {
	*Translational
	frame   int
	frames  []*frame
	animate bool
	repeat  bool
	render  bool
}

type frame struct {
	allocation       *util.ListNode
	wait, waitLength int
	translation      mgl32.Vec2
}

/*
Animations work as follows:
define each frame of the animation as a set of vertices and textures
then generate set of transitions between each set of frames
calling start, freeze and reset control the progression of these transitions
*/
func CreateAnimation(window *ggl.Window, texture string, size int32) (*Animation2D, error) {
	t, err := CreateTranslationalRenderer(window, texture, size)
	if err != nil {
		return nil, err
	}

	return &Animation2D{
		Translational: t,
	}, nil
}

func (a *Animation2D) AddFrame(verts, texs []mgl32.Vec2, transition util.FrameTransition) (int, error) {
	allocation, err := a.AllocateAndSetVertices(verts, texs)
	if err != nil {
		return -1, err
	}

	f := &frame{
		allocation:  allocation,
		wait:        0,
		waitLength:  transition.Delay,
		translation: transition.Translation,
	}

	a.frames = append(a.frames, f)

	return len(a.frames) - 1, nil
}

func (a *Animation2D) Start() {
	a.render = true
	a.animate = true
	a.frame = 0
}

func (a *Animation2D) Freeze() {
	a.animate = false
}

func (a *Animation2D) Resume() {
	a.animate = true
}

func (a *Animation2D) Render() {
	if a.render {
		a.VAO().PrepRender()

		// Render only the current frame
		f := a.frames[a.frame]
		gl.DrawArrays(gl.TRIANGLES, int32(f.allocation.Index()), int32(f.allocation.Size()))

		// nesting is rather mooreish
		if a.animate {
			f.wait++
			if f.wait == f.waitLength {
				a.frame++
				f.wait = 0

				if a.frame == len(a.frames) {
					a.frame = 0
					if !a.repeat {
						a.render = false
						a.animate = false
					}
				}
			}
		}
	}
}
