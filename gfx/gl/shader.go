package gl

import (
	"fmt"
	gl "github.com/go-gl/glow/gl"
	"strings"
)

type Uniform struct {
	Name     string
	Location int32
	Type     uint32
	Size     int32
}

type Attribute struct {
	Name     string
	Location int32
	Type     uint32
	Size     int32
}

type Shader struct {
	glProgramID uint32
	uniforms    []Uniform
	attribs     []Attribute
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

	glProgramID, err := createProgramAndLink(glVertexShader, glFragmentShader)
	if err != nil {
		return nil, err
	}

	gl.DetachShader(glProgramID, glVertexShader)
	gl.DetachShader(glProgramID, glFragmentShader)

	shader := &Shader{glProgramID: glProgramID}

	shader.updateUniforms()
	shader.updateAttributes()

	return shader, nil
}

func (shader *Shader) Destroy() {
	if shader.glProgramID != 0 {
		gl.DeleteProgram(shader.glProgramID)
	}
}

func (shader *Shader) Use() {
	gl.UseProgram(shader.glProgramID)
}

func (shader *Shader) Unuse() {
	gl.UseProgram(0)
}

func createProgramAndLink(glVertexShader uint32, glFragmentShader uint32) (uint32, error) {
	glProgramID := gl.CreateProgram()
	gl.AttachShader(glProgramID, glVertexShader)
	gl.AttachShader(glProgramID, glFragmentShader)
	gl.LinkProgram(glProgramID)

	var status int32
	gl.GetProgramiv(glProgramID, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		programLog := getProgramLog(glProgramID)
		return 0, fmt.Errorf("program linking failed: %v", programLog)
	}

	return glProgramID, nil
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

func getProgramLog(glProgramID uint32) string {
	var logLen int32
	gl.GetProgramiv(glProgramID, gl.INFO_LOG_LENGTH, &logLen)
	programLog := strings.Repeat("\x00", int(logLen+1))
	gl.GetProgramInfoLog(glProgramID, logLen, nil, gl.Str(programLog))

	return programLog
}

func (shader *Shader) updateUniforms() {
	var uniformCount int32
	gl.GetProgramiv(shader.glProgramID, gl.ACTIVE_UNIFORMS, &uniformCount)
	uniforms := make([]Uniform, uniformCount, uniformCount)

	var buf [1024]uint8
	var bufLen int32
	var i uint32
	for i = 0; i < uint32(uniformCount); i++ {
		gl.GetActiveUniform(shader.glProgramID, i, 1024, &bufLen, &uniforms[i].Size, &uniforms[i].Type, &buf[0])
		uniforms[i].Name = string(buf[:bufLen])
		uniforms[i].Location = gl.GetUniformLocation(shader.glProgramID, &buf[0])
		fmt.Println(uniforms[i].Location)
	}

	shader.uniforms = uniforms
}

func (shader *Shader) updateAttributes() {
	var attribCount int32
	gl.GetProgramiv(shader.glProgramID, gl.ACTIVE_ATTRIBUTES, &attribCount)
	attribs := make([]Attribute, attribCount, attribCount)

	var buf [1024]uint8
	var bufLen int32
	var i uint32
	for i = 0; i < uint32(attribCount); i++ {
		gl.GetActiveAttrib(shader.glProgramID, i, 1024, &bufLen, &attribs[i].Size, &attribs[i].Type, &buf[0])
		attribs[i].Name = string(buf[:bufLen])
		attribs[i].Location = gl.GetAttribLocation(shader.glProgramID, &buf[0])
		fmt.Println(attribs[i].Name)
		fmt.Println(attribs[i].Location)
	}

	shader.attribs = attribs
}

func getUniformLocation() {
}
