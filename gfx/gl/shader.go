package gl

import (
	"fmt"
	gl "github.com/go-gl/glow/gl"
	"github.com/go-gl/mathgl/mgl32"
	"strings"
)

type Uniform struct {
	Location int32
	Type     uint32
	Size     int32
}

type Attribute struct {
	Location int32
	Type     uint32
	Size     int32
}

type Shader struct {
	glProgramID uint32
	uniforms    map[string]Uniform
	attribs     map[string]Attribute
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
	uniforms := make(map[string]Uniform, uniformCount)

	var buf [1024]uint8
	var bufLen int32
	var uniformSize int32
	var uniformType uint32
	var i uint32
	for i = 0; i < uint32(uniformCount); i++ {
		gl.GetActiveUniform(shader.glProgramID, i, 1024, &bufLen, &uniformSize, &uniformType, &buf[0])
		uniforms[string(buf[:bufLen])] = Uniform{
			Location: gl.GetUniformLocation(shader.glProgramID, &buf[0]),
			Size:     uniformSize,
			Type:     uniformType,
		}
	}
	shader.uniforms = uniforms
}

func (shader *Shader) updateAttributes() {
	var attribCount int32
	gl.GetProgramiv(shader.glProgramID, gl.ACTIVE_ATTRIBUTES, &attribCount)
	attribs := make(map[string]Attribute, attribCount)

	var buf [1024]uint8
	var bufLen int32
	var attribSize int32
	var attribType uint32
	var i uint32
	for i = 0; i < uint32(attribCount); i++ {
		gl.GetActiveAttrib(shader.glProgramID, i, 1024, &bufLen, &attribSize, &attribType, &buf[0])
		attribs[string(buf[:bufLen])] = Attribute{
			Location: gl.GetAttribLocation(shader.glProgramID, &buf[0]),
			Size:     attribSize,
			Type:     attribType,
		}
	}
	shader.attribs = attribs
}

func (shader *Shader) SetUniformMatrix(name string, matrix *mgl32.Mat4) error {
	if uniform, ok := shader.uniforms[name]; ok {
		gl.UniformMatrix4fv(uniform.Location, 1, false, &matrix[0])
		return nil
	}

	return fmt.Errorf("invalid uniform variable: %s", name)
}
