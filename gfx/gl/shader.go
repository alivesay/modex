package gl

import (
	"fmt"
	"strings"

	gogl "github.com/go-gl/gl/all-core/gl"
	"github.com/go-gl/mathgl/mgl32"
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
	glVertexShader, err := compileShader(vertexSource, gogl.VERTEX_SHADER)
	if err != nil {
		return nil, err
	}
	defer gogl.DeleteShader(glVertexShader)

	glFragmentShader, err := compileShader(fragmentSource, gogl.FRAGMENT_SHADER)
	if err != nil {
		return nil, err
	}
	defer gogl.DeleteShader(glFragmentShader)

	glProgramID, err := createProgramAndLink(glVertexShader, glFragmentShader)
	if err != nil {
		return nil, err
	}

	gogl.DetachShader(glProgramID, glVertexShader)
	gogl.DetachShader(glProgramID, glFragmentShader)

	shader := &Shader{glProgramID: glProgramID}

	shader.updateUniforms()
	shader.updateAttributes()

	return shader, nil
}

func (shader *Shader) Destroy() {
	if shader.glProgramID != 0 {
		gogl.DeleteProgram(shader.glProgramID)
	}
}

func (shader *Shader) Use() {
	gogl.UseProgram(shader.glProgramID)
}

func (shader *Shader) Unuse() {
	gogl.UseProgram(0)
}

func createProgramAndLink(glVertexShader uint32, glFragmentShader uint32) (uint32, error) {
	glProgramID := gogl.CreateProgram()
	gogl.AttachShader(glProgramID, glVertexShader)
	gogl.AttachShader(glProgramID, glFragmentShader)
	gogl.LinkProgram(glProgramID)

	var status int32
	gogl.GetProgramiv(glProgramID, gogl.LINK_STATUS, &status)
	if status == gogl.FALSE {
		programLog := getProgramLog(glProgramID)
		return 0, fmt.Errorf("program linking failed: %v", programLog)
	}

	return glProgramID, nil
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	glShader := gogl.CreateShader(shaderType)

	csources, free := gogl.Strs(source)
	gogl.ShaderSource(glShader, 1, csources, nil)
	free()
	gogl.CompileShader(glShader)

	var status int32
	gogl.GetShaderiv(glShader, gogl.COMPILE_STATUS, &status)
	if status == gogl.FALSE {
		shaderLog := getShaderLog(glShader)
		return 0, fmt.Errorf("shader compilation failed: %v: %v", source, shaderLog)
	}

	return glShader, nil
}

func getShaderLog(glShader uint32) string {
	var logLen int32
	gogl.GetShaderiv(glShader, gogl.INFO_LOG_LENGTH, &logLen)
	shaderLog := strings.Repeat("\x00", int(logLen+1))
	gogl.GetShaderInfoLog(glShader, logLen, nil, gogl.Str(shaderLog))

	return shaderLog
}

func getProgramLog(glProgramID uint32) string {
	var logLen int32
	gogl.GetProgramiv(glProgramID, gogl.INFO_LOG_LENGTH, &logLen)
	programLog := strings.Repeat("\x00", int(logLen+1))
	gogl.GetProgramInfoLog(glProgramID, logLen, nil, gogl.Str(programLog))

	return programLog
}

func (shader *Shader) updateUniforms() {
	var uniformCount int32
	gogl.GetProgramiv(shader.glProgramID, gogl.ACTIVE_UNIFORMS, &uniformCount)
	uniforms := make(map[string]Uniform, uniformCount)

	var buf [1024]uint8
	var bufLen int32
	var uniformSize int32
	var uniformType uint32
	var i uint32
	for i = 0; i < uint32(uniformCount); i++ {
		gogl.GetActiveUniform(shader.glProgramID, i, 1024, &bufLen, &uniformSize, &uniformType, &buf[0])
		uniforms[string(buf[:bufLen])] = Uniform{
			Location: gogl.GetUniformLocation(shader.glProgramID, &buf[0]),
			Size:     uniformSize,
			Type:     uniformType,
		}
	}
	shader.uniforms = uniforms
}

func (shader *Shader) updateAttributes() {
	var attribCount int32
	gogl.GetProgramiv(shader.glProgramID, gogl.ACTIVE_ATTRIBUTES, &attribCount)
	attribs := make(map[string]Attribute, attribCount)

	var buf [1024]uint8
	var bufLen int32
	var attribSize int32
	var attribType uint32
	var i uint32
	for i = 0; i < uint32(attribCount); i++ {
		gogl.GetActiveAttrib(shader.glProgramID, i, 1024, &bufLen, &attribSize, &attribType, &buf[0])
		attribs[string(buf[:bufLen])] = Attribute{
			Location: gogl.GetAttribLocation(shader.glProgramID, &buf[0]),
			Size:     attribSize,
			Type:     attribType,
		}
	}
	shader.attribs = attribs
}

func (shader *Shader) SetUniformMatrix(name string, matrix *mgl32.Mat4) error {
	if uniform, ok := shader.uniforms[name]; ok {
		gogl.UniformMatrix4fv(uniform.Location, 1, false, &matrix[0])
		return nil
	}

	return fmt.Errorf("invalid uniform variable: %s", name)
}
