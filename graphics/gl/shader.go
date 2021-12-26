package gl

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
	Id uint32
}

func CreateProgram(Id uint32) *Program {
	if Id == 0 {
		Id = gl.CreateProgram()
	}

	return &Program{Id}
}

func (p *Program) AttachShader(s *shader) {
	gl.AttachShader(p.Id, s.id)
}

// Cache already fetched shaders
var (
	loadedShaders = map[string]*shader{}
)

func (p *Program) loadShader(loc string, shaderType uint32) error {
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

func (p *Program) Use() {
	gl.UseProgram(p.Id)
}

func (p *Program) Release() {
	gl.UseProgram(0)
}

func (p *Program) Link() {
	gl.LinkProgram(p.Id)
}
