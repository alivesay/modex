package gfx

import (
	"github.com/alivesay/modex/core"
	"github.com/alivesay/modex/gfx/gl"
)

type VideoMode struct {
	Width      uint16
	Height     uint16
	Bpp        uint8
	Fullscreen bool
}

type Video struct {
	osWindow   *OSWindow
	glState    *gl.State
	glShader   *gl.Shader
	glRenderer *gl.Renderer
}

const initialTitle string = "modex"
const initialWidth uint16 = 320
const initialHeight uint16 = 240

const defaultVertexShaderGLSL string = `
#version 100
precision highp float;
uniform mat4 uProjectionMatrix;
attribute vec2 Position;
void main() {
	gl_Position = uProjectionMatrix * vec4(Position, 0.0, 1.0);
}` + "\x00"

const defaultFragmentShaderGLSL string = `
#version 100
precision highp float;
vec4 color = vec4(1.0, 1.0, 1.0, 1.0);
void main() {
	if (mod(gl_FragCoord.y, 2.0f) < 1.0) {
		gl_FragColor = color;
	} else {
		gl_FragColor = vec4(0.0, 0.0, 0.0, 0.0);
	}
}` + "\x00"

func NewVideo() (*Video, error) {

	osWindow, err := NewOSWindow(initialTitle, initialWidth, initialHeight)
	if err != nil {
		return nil, err
	}

	glState := &gl.State{}
	glState.Init()

	glShader, err := gl.NewShader(defaultVertexShaderGLSL, defaultFragmentShaderGLSL)
	if err != nil {
		return nil, err
	}

	if err := osWindow.SetViewport2D(); err != nil {
		return nil, err
	}

	osWindow.SetShader(glShader)

	return &Video{
		osWindow:   osWindow,
		glState:    glState,
		glShader:   glShader,
		glRenderer: gl.NewRenderer(),
	}, nil
}

func (video *Video) Destroy() {
	video.osWindow.Destroy()
}

func (video *Video) SetMode(mode *VideoMode) {
}

func (video *Video) Render() {
	//	video.glRenderer.Render()
	video.osWindow.Update()
	video.osWindow.Render()

	if err := gl.GetGLError(); err != nil {
		core.Log(core.LOG_ERR, err)
	}
}
