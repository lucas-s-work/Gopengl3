package shader

import (
	"fmt"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/lucas-s-work/gopengl3/util"
)

type shader struct {
	id uint32
}

type Program struct {
	Id         uint32
	uniforms   map[string]*Uniform
	attributes map[string]uint32
}

func CreateProgram(Id uint32) *Program {
	if Id == 0 {
		Id = gl.CreateProgram()
	}

	return &Program{
		Id:         Id,
		uniforms:   make(map[string]*Uniform),
		attributes: make(map[string]uint32),
	}
}

func (p *Program) AttachShader(s *shader) {
	gl.AttachShader(p.Id, s.id)
}

// Cache already fetched shaders
var (
	loadedShaders = map[string]*shader{}
)

func (p *Program) LoadShader(loc string, shaderType uint32) error {
	// If shader reused and already loaded then just use that
	if s, ok := loadedShaders[loc]; ok {
		p.AttachShader(s)

		return nil
	}

	shaderString, err := util.ReadFile(loc)

	if err != nil {
		return err
	}

	// Will cause errors if we don't add empty char to end of shader source
	shaderString += "\x00"

	shaderId := gl.CreateShader(shaderType)
	source, free := gl.Strs(shaderString)

	gl.ShaderSource(shaderId, 1, source, nil)
	free()
	gl.CompileShader(shaderId)

	// Fetch debug information about compilation errors
	var status int32
	gl.GetShaderiv(shaderId, gl.COMPILE_STATUS, &status)

	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shaderId, gl.INFO_LOG_LENGTH, &logLength)
		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shaderId, logLength, nil, gl.Str(log))
		return fmt.Errorf("Failed to compile shader: %s, %s", loc, log)
	}

	loadedShaders[loc] = &shader{shaderId}
	gl.AttachShader(p.Id, shaderId)

	return nil
}

func (p *Program) Use() error {
	if p.Id == 0 {
		return fmt.Errorf("Cannot use deleted program")
	}

	gl.UseProgram(p.Id)

	return nil
}

func (p *Program) Release() {
	gl.UseProgram(0)
}

func (p *Program) Link() error {
	if p.Id == 0 {
		return fmt.Errorf("Cannot link deleted program")
	}

	gl.LinkProgram(p.Id)

	return nil
}

func (p *Program) Delete() {
	gl.DeleteProgram(p.Id)
	p.Id = 0
}

func (p *Program) AttachUniform(name string, value interface{}) error {
	if _, ok := p.uniforms[name]; ok {
		return fmt.Errorf("Uniform with name: %s already attached", name)
	}

	u := Uniform{
		gl.GetUniformLocation(p.Id, gl.Str(name+"\x00")),
		value,
		true,
	}

	p.uniforms[name] = &u

	p.Use()
	if err := u.Update(); err != nil {
		return err
	}
	p.Release()

	return nil
}

func (p *Program) UpdateUniforms() error {
	p.Use()
	for _, u := range p.uniforms {
		if err := u.Update(); err != nil {
			return err
		}
	}
	p.Release()

	return nil
}

func (p *Program) AttachAttribute(n string, size int32) error {
	attribRaw := gl.GetAttribLocation(p.Id, gl.Str(n+"\x00"))

	if attribRaw == -1 {
		return fmt.Errorf("Unable to find attribute: %s", n)
	}

	attrib := uint32(attribRaw)

	p.attributes[n] = attrib

	if err := p.EnableAttribute(n); err != nil {
		return err
	}
	gl.VertexAttribPointer(attrib, size, gl.FLOAT, false, 0, nil)

	return nil
}

func (p *Program) EnableAttribute(n string) error {
	if attrib, ok := p.attributes[n]; ok {
		gl.EnableVertexAttribArray(attrib)
	} else {
		return fmt.Errorf("No attribute: %s registered", n)
	}

	return nil
}

func (p *Program) DisableAttribute(n string) error {
	if attrib, ok := p.attributes[n]; ok {
		gl.DisableVertexAttribArray(attrib)
	} else {
		return fmt.Errorf("No attribute: %s registered", n)
	}

	return nil
}

func (p *Program) EnableAttributes() {
	for _, attrib := range p.attributes {
		gl.EnableVertexAttribArray(attrib)
	}
}

func (p *Program) DisableAttributes() {
	for _, attrib := range p.attributes {
		gl.DisableVertexAttribArray(attrib)
	}
}
