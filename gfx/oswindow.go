package gfx

import (
	"image"
	_ "image/draw"
	_ "image/png"
	"math/rand"
	"os"

	"github.com/alivesay/modex/core"
	"github.com/alivesay/modex/events"
	"github.com/alivesay/modex/gfx/gl"

	"github.com/go-gl/glfw/v3.1/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

// OSWindow manages the creation of new OpenGL client windows.
type OSWindow struct {
	glfwWindow       *glfw.Window
	keyEventCallback events.KeyCallback
	viewport         *gl.Viewport
	shader           *gl.Shader
	mesh             *gl.Mesh
	texture          *gl.Texture
	fps              *FPS
}

// NewOSWindow creates a new OSWindow.
func NewOSWindow(title string, w uint16, h uint16) (*OSWindow, error) {
	glfw.WindowHint(glfw.ClientAPI, glfw.OpenGLESAPI)
	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 0)
	//glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
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
		keyEventCallback: defaultKeyCallback,
		fps:              NewFPS(),
	}

	glfwWindow.SetFramebufferSizeCallback(window.framebufferResizeCallback)
	glfwWindow.SetKeyCallback(window.keyCallback)

	return window, nil
}

// Destroy implements an OSWindow destructor.
func (window *OSWindow) Destroy() {
	window.mesh.Destroy()
	window.glfwWindow.Destroy()
}

// TODO: should be abstracted to something like Use2DMode()

// SetViewportPerspective maximizes the OSWindow's viewport and sets up
// a default perspective projection matrix.
func (window *OSWindow) SetViewportPerspective() error {
	vw, vh := window.glfwWindow.GetFramebufferSize()
	window.viewport = gl.NewViewport(0, 0, vw, vh)

	if err := window.viewport.SetPerspective(); err != nil {
		return err
	}

	return nil
}

// SetViewport2D maximizes the OSWindow's viewport and sets up
// a default orthographic projection matrix.
func (window *OSWindow) SetViewport2D() error {

	vw, vh := window.glfwWindow.GetFramebufferSize()
	window.viewport = gl.NewViewport(0, 0, vw, vh)

	if err := window.viewport.SetOrtho2D(); err != nil {
		return err
	}
	window.LoadTexture()
	return nil
}

// SetShader sets the root Shader to use for the window's rendering.
func (window *OSWindow) SetShader(shader *gl.Shader) {
	window.shader = shader
}

// Update initiates and propogates the application update phase.
func (window *OSWindow) Update() {
	events.Poll()
	if window.glfwWindow.ShouldClose() {
		core.GetInstanceApplication().ShutdownRequested = true
	}
	window.DrawRandomTriangles()
}

// Render initiates and propogates the application render phase.
func (window *OSWindow) Render() {
	// this shader will go in an overlay or target fbo
	window.viewport.Render()

	// TODO: why bind via mesh?
	window.mesh.VBO.Bind()
	window.texture.Bind()
	if window.shader != nil {
		window.shader.Use()
		//viewMat := mgl32.LookAt(60, 40, -32, 320, 0, 50, 0, 1, 0)
		//		viewMat := mgl32.LookAt(0, 32, 260, 200, 0, 0, 0, 1, 0)
		//viewMat := mgl32.LookAt(0, 20, 32, 0, 0, 0, 0, 1, 0)
		viewMat := mgl32.Ident4()
		modelMat := mgl32.Ident4()

		// TODO:
		// these should be set automatically
		window.shader.SetUniformMatrix("uModelMatrix", &modelMat)
		window.shader.SetUniformMatrix("uViewMatrix", &viewMat)
		window.shader.SetUniformMatrix("uProjectionMatrix", window.viewport.ProjMat)
		window.shader.SetUniformi("texture1", 0)
	}

	window.mesh.VBO.Render()
	window.swap()
	window.texture.Unbind()
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

// TODO:

func defaultKeyCallback(keyEvent *events.KeyEvent) {
	core.Log(core.LogDebug, keyEvent)
}

func (window *OSWindow) LoadTexture() {
	infile, err := os.Open("helm.png")
	if err != nil {
		panic(err)
	}
	defer infile.Close()

	src, _, err := image.Decode(infile)
	if err != nil {
		panic(err)
	}

	texParams := []gl.TextureParameter{
		gl.TextureParameter{Name: gl.TextureMinFilter, Value: gl.Nearest},
		gl.TextureParameter{Name: gl.TextureMagFilter, Value: gl.Nearest},
		gl.TextureParameter{Name: gl.TextureWrapS, Value: gl.ClampToEdge},
		gl.TextureParameter{Name: gl.TextureWrapT, Value: gl.ClampToEdge},
	}
	window.texture, err = gl.NewTextureFromImage(src, gl.Texture2D, texParams)
	// TODO: replace these with something like:
	// texture.SetMinFilter(), tex.GenerateMipMap()
	if err != nil {
		panic(err)
	}
}

func (window *OSWindow) SetupTestMesh() {
	var err error

	window.mesh, err = gl.NewMesh(gl.GLTriangles, gl.DynamicDraw, 2048)
	if err != nil {
		panic(err)
	}
}

func (window *OSWindow) DrawRandomTriangles() {
	window.mesh.ClearBuffer()

	for i := 0; i < 1000; i++ {
		x := rand.Float32() * float32(window.viewport.Rect.Dx())
		y := rand.Float32() * float32(window.viewport.Rect.Dy())

		window.mesh.AddVertex(gl.Vertex{x, y + 32, 0.0, 0.0, 1.0})      // bottom left
		window.mesh.AddVertex(gl.Vertex{x + 32, y + 32, 0.0, 1.0, 1.0}) // bottom right
		window.mesh.AddVertex(gl.Vertex{x, y, 0.0, 0.0, 0.0})           // top left

		window.mesh.AddVertex(gl.Vertex{x + 32, y, 0.0, 1.0, 0.0})      // top right
		window.mesh.AddVertex(gl.Vertex{x, y, 0.0, 0.0, 0.0})           // top left
		window.mesh.AddVertex(gl.Vertex{x + 32, y + 32, 0.0, 1.0, 1.0}) // bottom right
	}

	window.mesh.SyncBuffer()
}
