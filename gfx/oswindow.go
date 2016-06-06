package gfx

import (
	"github.com/alivesay/modex/core"
	"github.com/alivesay/modex/events"
	"github.com/go-gl/glfw/v3.1/glfw"
)

type OSWindow struct {
	glfwWindow       *glfw.Window
	keyEventCallback events.KeyCallback
}

func NewOSWindow(title string, w uint16, h uint16) (*OSWindow, error) {
	glfw.WindowHint(glfw.ClientAPI, glfw.OpenGLESAPI)
	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.DoubleBuffer, glfw.True)

	glfwWindow, err := glfw.CreateWindow(int(w), int(h), title, nil, nil)
	if err != nil {
		return nil, err
	}

	glfwWindow.MakeContextCurrent()

	window := &OSWindow{
		glfwWindow:       glfwWindow,
		keyEventCallback: DefaultKeyCallback,
	}

	glfwWindow.SetKeyCallback(window.keyCallback)

	return window, nil
}

func (window *OSWindow) Destroy() {
	window.glfwWindow.Destroy()
}

func (window *OSWindow) Swap() {
	window.glfwWindow.SwapBuffers()
	events.Poll()
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
