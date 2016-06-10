package gfx

import (
	"github.com/alivesay/modex/core"
	"github.com/alivesay/modex/events"
	"github.com/alivesay/modex/gfx/gl"
	"github.com/alivesay/modex/gfx/pixbuf"
	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

type OSWindow struct {
	glfwWindow       *glfw.Window
	keyEventCallback events.KeyCallback
	viewport         *gl.Viewport
	ProjMat          mgl32.Mat4
	shader           *gl.Shader
	BgColor          *pixbuf.RGBA32
}

func NewOSWindow(title string, w uint16, h uint16) (*OSWindow, error) {
	glfw.WindowHint(glfw.ClientAPI, glfw.OpenGLESAPI)
	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 0)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.DoubleBuffer, glfw.True)

	glfwWindow, err := glfw.CreateWindow(int(w), int(h), title, nil, nil)
	if err != nil {
		return nil, err
	}

	glfwWindow.MakeContextCurrent()

	window := &OSWindow{
		glfwWindow:       glfwWindow,
		keyEventCallback: DefaultKeyCallback,
		BgColor:          &pixbuf.RGBA32{Packed: 0x3366CCFF},
	}

	glfwWindow.SetKeyCallback(window.keyCallback)

	return window, nil
}

func (window *OSWindow) Destroy() {
	window.glfwWindow.Destroy()
}

func (window *OSWindow) SetViewport2D() error {
	vw, vh := window.glfwWindow.GetFramebufferSize()
	window.viewport = &gl.Viewport{0, 0, int32(vw), int32(vh), &window.ProjMat}

	if err := window.viewport.SetOrtho2D(); err != nil {
		return err
	}

	return nil
}

func (window *OSWindow) SetShader(shader *gl.Shader) {
	window.shader = shader
}

func (window *OSWindow) Update() {
	events.Poll()
}

func (window *OSWindow) Render() {
	window.viewport.Clear(window.BgColor)
	// this shader will go in an overlay or target fbo
	if window.shader != nil {
		window.shader.Use()
	}

	window.swap()

	if window.shader != nil {
		window.shader.Unuse()
	}
}

func (window *OSWindow) swap() {
	window.glfwWindow.SwapBuffers()
	if window.glfwWindow.ShouldClose() {
		core.GetInstanceApplication().ShutdownRequested = true
	}
}

func (window *OSWindow) keyCallback(glfwWindow *glfw.Window, key glfw.Key, scancode int, action glfw.Action, modifierKey glfw.ModifierKey) {
	if window.keyEventCallback != nil {
		window.keyEventCallback(&events.KeyEvent{
			Window:   glfwWindow,
			Key:      events.Key(key),
			Scancode: scancode,
			Action:   events.Action(action),
			Mods:     events.ModifierKey(modifierKey),
		})
	}
}

func DefaultKeyCallback(keyEvent *events.KeyEvent) {
	core.Log(core.LOG_DEBUG, keyEvent)
}
