package graphics

import (
	"fmt"
	"sort"

	"github.com/go-gl/gl/v4.1-core/gl"
	ggl "github.com/lucas-s-work/gopengl3/graphics/gl"
)

const (
	maxJobs = 256 // Maximum number of render jobs queued per frame
)

type Context struct {
	renderers map[int][]Renderer
	jobs      chan func()
	layers    []int
	sync      chan struct{}
	window    *ggl.Window
	useSync   bool
}

func CreateContext(window *ggl.Window) *Context {
	ctx := &Context{
		renderers: make(map[int][]Renderer),
		jobs:      make(chan func(), maxJobs),
		layers:    []int{},
		window:    window,
		sync:      make(chan struct{}),
	}

	return ctx
}

func (ctx *Context) Attach(renderer Renderer, layer int) {
	layerExist := false
	for _, l := range ctx.layers {
		if l == layer {
			layerExist = true
		}

		if l > layer {
			break
		}
	}

	renderer.SetLayer(layer)

	if !layerExist {
		ctx.layers = append(ctx.layers, layer)
		renderer.SetId(0)
		ctx.renderers[layer] = []Renderer{renderer}
	} else {
		indexFound := false
		for i, r := range ctx.renderers[layer] {
			if r == nil {
				renderer.SetId(i)
				ctx.renderers[layer][i] = renderer
				indexFound = true

				break
			}
		}

		if !indexFound {
			renderer.SetId(len(ctx.renderers[layer]))
			ctx.renderers[layer] = append(ctx.renderers[layer], renderer)
		}
	}

	sort.Ints(ctx.layers)
}

func (ctx *Context) Detach(renderer Renderer) {
	ctx.renderers[renderer.GetLayer()][renderer.GetId()] = nil
	renderer.Delete()
}

func (ctx *Context) GetSync() chan<- struct{} {
	ctx.useSync = true
	return ctx.sync
}

func (ctx *Context) Delete() {
	for _, l := range ctx.layers {
		for _, r := range ctx.renderers[l] {
			r.Delete()
		}
	}

	ctx.renderers = nil
	ctx.layers = nil
	ctx.jobs = nil
}

func (ctx *Context) Window() *ggl.Window {
	return ctx.window
}

func (ctx *Context) AddJob(job func()) error {
	select {
	case ctx.jobs <- job:
		return nil
	default:
		return fmt.Errorf("Unable to place render job, queue size: %v full", maxJobs)
	}
}

func (ctx *Context) executeJobs() {
	if ctx.useSync {
		for {
			shouldReturn := false
		loop:
			for {
				select {
				case j := <-ctx.jobs:
					j()
				default:
					shouldReturn = true
					break loop
				}
			}

			// We wait for the context sync *After* performing the jobs
			// To allow for tick synchronization without nil pointers
			// I.e Tick() will block on these jobs being complete
			// before executing anything, this allows any initialization logic
			// To occur pre tick
			<-ctx.sync
			if shouldReturn {
				return
			}
		}
	}
}

func (ctx *Context) Render() {
	gl.ClearColor(0.0, 0.0, 0.0, 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT)

	ctx.executeJobs()

	for _, l := range ctx.layers {
		for _, r := range ctx.renderers[l] {
			if r != nil {
				r.Render()
			}
		}
	}

	ctx.window.SwapBuffers()
	ctx.window.PollInput()
}
