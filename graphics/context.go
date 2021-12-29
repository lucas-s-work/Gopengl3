package graphics

import "sort"

type Context struct {
	renderers map[int][]Renderer
	jobs      []func()
	layers    []int
	sync      chan struct{}
	useSync   bool
}

func CreateContext() *Context {
	ctx := &Context{
		renderers: make(map[int][]Renderer),
		jobs:      []func(){},
		layers:    []int{},
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
	ctx.executeJobs()

	for _, l := range ctx.layers {
		for _, r := range ctx.renderers[l] {
			r.Render()
		}
	}
}
