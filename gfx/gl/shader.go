package gl

import (
	"fmt"
	gl "github.com/go-gl/glow/gl"
	"strings"
)

type Shader struct {
	glProgram uint32
}

func NewShader(vertexSource string, fragmentSource string) (*Shader, error) {
	glVertexShader, err := compileShader(vertexSource, gl.VERTEX_SHADER)
	if err != nil {
		return nil, err
	}
	defer gl.DeleteShader(glVertexShader)

	glFragmentShader, err := compileShader(fragmentSource, gl.FRAGMENT_SHADER)
	if err != nil {
		return nil, err
	}
	defer gl.DeleteShader(glFragmentShader)

	glProgram, err := createProgramAndLink(glVertexShader, glFragmentShader)
	if err != nil {
		return nil, err
	}

	gl.DetachShader(glProgram, glVertexShader)
	gl.DetachShader(glProgram, glFragmentShader)

	return &Shader{glProgram: glProgram}, nil
}

func (shader *Shader) Destroy() {
	if shader.glProgram != 0 {
		gl.DeleteProgram(shader.glProgram)
	}
}

func (shader *Shader) Use() {
	gl.UseProgram(shader.glProgram)
}

func (shader *Shader) Unuse() {
	gl.UseProgram(0)
}

func createProgramAndLink(glVertexShader uint32, glFragmentShader uint32) (uint32, error) {
	glProgram := gl.CreateProgram()
	gl.AttachShader(glProgram, glVertexShader)
	gl.AttachShader(glProgram, glFragmentShader)
	gl.LinkProgram(glProgram)

	var status int32
	gl.GetProgramiv(glProgram, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		programLog := getProgramLog(glProgram)
		return 0, fmt.Errorf("program linking failed: %v", programLog)
	}

	return glProgram, nil
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	glShader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(glShader, 1, csources, nil)
	free()
	gl.CompileShader(glShader)

	var status int32
	gl.GetShaderiv(glShader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		shaderLog := getShaderLog(glShader)
		return 0, fmt.Errorf("shader compilation failed: %v: %v", source, shaderLog)
	}

	return glShader, nil
}

func getShaderLog(glShader uint32) string {
	var logLen int32
	gl.GetShaderiv(glShader, gl.INFO_LOG_LENGTH, &logLen)
	shaderLog := strings.Repeat("\x00", int(logLen+1))
	gl.GetShaderInfoLog(glShader, logLen, nil, gl.Str(shaderLog))

	return shaderLog
}

func getProgramLog(glProgram uint32) string {
	var logLen int32
	gl.GetProgramiv(glProgram, gl.INFO_LOG_LENGTH, &logLen)
	programLog := strings.Repeat("\x00", int(logLen+1))
	gl.GetProgramInfoLog(glProgram, logLen, nil, gl.Str(programLog))

	return programLog
}
