package gl

import (
	"fmt"
	gles2 "github.com/go-gl/gl/v3.1/gles2"
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
	glVertexShader, err := compileShader(vertexSource, gles2.VERTEX_SHADER)
	if err != nil {
		return nil, err
	}
	defer gles2.DeleteShader(glVertexShader)

	glFragmentShader, err := compileShader(fragmentSource, gles2.FRAGMENT_SHADER)
	if err != nil {
		return nil, err
	}
	defer gles2.DeleteShader(glFragmentShader)

	glProgramID, err := createProgramAndLink(glVertexShader, glFragmentShader)
	if err != nil {
		return nil, err
	}

	gles2.DetachShader(glProgramID, glVertexShader)
	gles2.DetachShader(glProgramID, glFragmentShader)

	shader := &Shader{glProgramID: glProgramID}

	shader.updateUniforms()
	shader.updateAttributes()

	return shader, nil
}

func (shader *Shader) Destroy() {
	if shader.glProgramID != 0 {
		gles2.DeleteProgram(shader.glProgramID)
	}
}

func (shader *Shader) Use() {
	gles2.UseProgram(shader.glProgramID)
}

func (shader *Shader) Unuse() {
	gles2.UseProgram(0)
}

func createProgramAndLink(glVertexShader uint32, glFragmentShader uint32) (uint32, error) {
	glProgramID := gles2.CreateProgram()
	gles2.AttachShader(glProgramID, glVertexShader)
	gles2.AttachShader(glProgramID, glFragmentShader)
	gles2.LinkProgram(glProgramID)

	var status int32
	gles2.GetProgramiv(glProgramID, gles2.LINK_STATUS, &status)
	if status == gles2.FALSE {
		programLog := getProgramLog(glProgramID)
		return 0, fmt.Errorf("program linking failed: %v", programLog)
	}

	return glProgramID, nil
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	glShader := gles2.CreateShader(shaderType)

	csources, free := gles2.Strs(source)
	gles2.ShaderSource(glShader, 1, csources, nil)
	free()
	gles2.CompileShader(glShader)

	var status int32
	gles2.GetShaderiv(glShader, gles2.COMPILE_STATUS, &status)
	if status == gles2.FALSE {
		shaderLog := getShaderLog(glShader)
		return 0, fmt.Errorf("shader compilation failed: %v: %v", source, shaderLog)
	}

	return glShader, nil
}

func getShaderLog(glShader uint32) string {
	var logLen int32
	gles2.GetShaderiv(glShader, gles2.INFO_LOG_LENGTH, &logLen)
	shaderLog := strings.Repeat("\x00", int(logLen+1))
	gles2.GetShaderInfoLog(glShader, logLen, nil, gles2.Str(shaderLog))

	return shaderLog
}

func getProgramLog(glProgramID uint32) string {
	var logLen int32
	gles2.GetProgramiv(glProgramID, gles2.INFO_LOG_LENGTH, &logLen)
	programLog := strings.Repeat("\x00", int(logLen+1))
	gles2.GetProgramInfoLog(glProgramID, logLen, nil, gles2.Str(programLog))

	return programLog
}

func (shader *Shader) updateUniforms() {
	var uniformCount int32
	gles2.GetProgramiv(shader.glProgramID, gles2.ACTIVE_UNIFORMS, &uniformCount)
	uniforms := make(map[string]Uniform, uniformCount)

	var buf [1024]uint8
	var bufLen int32
	var uniformSize int32
	var uniformType uint32
	var i uint32
	for i = 0; i < uint32(uniformCount); i++ {
		gles2.GetActiveUniform(shader.glProgramID, i, 1024, &bufLen, &uniformSize, &uniformType, &buf[0])
		uniforms[string(buf[:bufLen])] = Uniform{
			Location: gles2.GetUniformLocation(shader.glProgramID, &buf[0]),
			Size:     uniformSize,
			Type:     uniformType,
		}
	}
	shader.uniforms = uniforms
}

func (shader *Shader) updateAttributes() {
	var attribCount int32
	gles2.GetProgramiv(shader.glProgramID, gles2.ACTIVE_ATTRIBUTES, &attribCount)
	attribs := make(map[string]Attribute, attribCount)

	var buf [1024]uint8
	var bufLen int32
	var attribSize int32
	var attribType uint32
	var i uint32
	for i = 0; i < uint32(attribCount); i++ {
		gles2.GetActiveAttrib(shader.glProgramID, i, 1024, &bufLen, &attribSize, &attribType, &buf[0])
		attribs[string(buf[:bufLen])] = Attribute{
			Location: gles2.GetAttribLocation(shader.glProgramID, &buf[0]),
			Size:     attribSize,
			Type:     attribType,
		}
	}
	shader.attribs = attribs
}

func (shader *Shader) SetUniformMatrix(name string, matrix *mgl32.Mat4) error {
	if uniform, ok := shader.uniforms[name]; ok {
		gles2.UniformMatrix4fv(uniform.Location, 1, false, &matrix[0])
		return nil
	}

	return fmt.Errorf("invalid uniform variable: %s", name)
}
