package graphics

import (
	"sort"

	"github.com/go-gl/gl/v4.1-core/gl"
	ggl "github.com/lucas-s-work/gopengl3/graphics/gl"
)

type Context struct {
	renderers map[int][]Renderer
	jobs      []func()
	layers    []int
	sync      chan struct{}
	window    *ggl.Window
	useSync   bool
}

func CreateContext(window *ggl.Window) *Context {
	ctx := &Context{
		renderers: make(map[int][]Renderer),
		jobs:      []func(){},
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

	if !layerExist {
		ctx.layers = append(ctx.layers, layer)
		ctx.renderers[layer] = []Renderer{renderer}
	} else {
		ctx.renderers[layer] = append(ctx.renderers[layer], renderer)
	}

	sort.Ints(ctx.layers)
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

	ctx.renderers = map[int][]Renderer{}
	ctx.layers = []int{}
	ctx.jobs = []func(){}
}

func (ctx *Context) AddJob(job func()) {
	ctx.jobs = append(ctx.jobs, job)
}

func (ctx *Context) executeJobs() {
	if ctx.useSync {
		<-ctx.sync
		for _, j := range ctx.jobs {
			j()
		}
		ctx.jobs = []func(){}
	}
}

func (ctx *Context) Render() {
	gl.ClearColor(0.0, 0.0, 0.0, 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT)

	ctx.executeJobs()

	for _, l := range ctx.layers {
		for _, r := range ctx.renderers[l] {
			r.Render()
		}
	}

	ctx.window.SwapBuffers()
	ctx.window.PollInput()
}
