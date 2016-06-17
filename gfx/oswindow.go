package gfx

import (
	"math/rand"

	"github.com/alivesay/modex/core"
	"github.com/alivesay/modex/events"
	"github.com/alivesay/modex/gfx/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

type OSWindow struct {
	glfwWindow       *glfw.Window
	keyEventCallback events.KeyCallback
	viewport         *gl.Viewport
	shader           *gl.Shader
	mesh             *gl.Mesh
	fps              *FPS
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
	// TODO
	glfw.SwapInterval(0)

	window := &OSWindow{
		glfwWindow:       glfwWindow,
		keyEventCallback: DefaultKeyCallback,
		fps:              NewFPS(),
	}

	glfwWindow.SetFramebufferSizeCallback(window.framebufferResizeCallback)
	glfwWindow.SetKeyCallback(window.keyCallback)

	return window, nil
}

func (window *OSWindow) Destroy() {
	window.mesh.Destroy()
	window.glfwWindow.Destroy()
}

func (window *OSWindow) SetViewportPerspective() error {
	vw, vh := window.glfwWindow.GetFramebufferSize()
	window.viewport = gl.NewViewport(0, 0, int32(vw), int32(vh))

	if err := window.viewport.SetPerspective(); err != nil {
		return err
	}

	return nil
}

// TODO: should be abstracted to something like Use2DMode()
func (window *OSWindow) SetViewport2D() error {
	vw, vh := window.glfwWindow.GetFramebufferSize()
	window.viewport = gl.NewViewport(0, 0, int32(vw), int32(vh))

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
	if window.glfwWindow.ShouldClose() {
		core.GetInstanceApplication().ShutdownRequested = true
	}
	window.DrawRandomTriangles()
}

func (window *OSWindow) Render() {
	// this shader will go in an overlay or target fbo
	window.viewport.Render()

	window.mesh.VBO.Bind()
	if window.shader != nil {
		window.shader.Use()
		//viewMat := mgl32.LookAt(60, 40, -10, 100, 0, 50, 0, 1, 0)
		//		viewMat := mgl32.LookAt(0, 30, 260, 200, 0, 0, 0, 1, 0)
		//viewMat := mgl32.LookAt(0, 20, 10, 0, 0, 0, 0, 1, 0)
		viewMat := mgl32.Ident4()
		modelMat := mgl32.Ident4()
		// TODO:
		// these should be set automatically
		window.shader.SetUniformMatrix("uModelMatrix", &modelMat)
		window.shader.SetUniformMatrix("uViewMatrix", &viewMat)
		window.shader.SetUniformMatrix("uProjectionMatrix", window.viewport.ProjMat)
	}
	window.mesh.VBO.Render()
	window.swap()

	window.mesh.VBO.Unbind()

	if window.shader != nil {
		window.shader.Unuse()
	}

	window.fps.Update(window.glfwWindow)
}

func (window *OSWindow) swap() {
	window.glfwWindow.SwapBuffers()
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

func (window *OSWindow) framebufferResizeCallback(glfwWindow *glfw.Window, width, height int) {
	window.SetViewport2D()
}

func DefaultKeyCallback(keyEvent *events.KeyEvent) {
	core.Log(core.LogDebug, keyEvent)
}

func (window *OSWindow) SetupTestMesh() {
	var err error
	window.mesh, err = gl.NewMesh(gl.StaticDraw, 2048)
	if err != nil {
		panic(err)
	}

	err = window.mesh.AddVertexAttrib(gl.VertexAttrib{0, 3, gl.GLFloat, false, 0, 0})
	if err != nil {
		panic(err)
	}
}

func (window *OSWindow) DrawRandomTriangles() {
	window.mesh.ClearBuffer()

	for i := 0; i < 10000; i++ {
		x := rand.Float32() * float32(window.viewport.W)
		y := rand.Float32() * float32(window.viewport.H)
		window.mesh.AddVertex(gl.Vertex{x, y, 0.0})
		window.mesh.AddVertex(gl.Vertex{x + 10, y + 10, 0.0})
		window.mesh.AddVertex(gl.Vertex{x - 10, y + 10, 0.0})
	}

	window.mesh.SyncBuffer()
}
