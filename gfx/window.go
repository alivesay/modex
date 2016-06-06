package gfx

import (
	"github.com/go-gl/glfw/v3.1/glfw"
)

type OSWindow struct {
	glfwWindow *glfw.Window
}

func NewOSWindow(title string, w uint16, h uint16) (*OSWindow, error) {

	glfw.WindowHint(glfw.ClientAPI, glfw.OpenGLESAPI)
	glfw.WindowHint(glfw.ContextVersionMajor, 2)

	glfwWindow, err := glfw.CreateWindow(int(w), int(h), title, nil, nil)
	if err != nil {
		return nil, err
	}

	glfwWindow.MakeContextCurrent()

	return &OSWindow{glfwWindow: glfwWindow}, nil
}

func (window *OSWindow) Destroy() {
	window.glfwWindow.Destroy()
}

func (window *OSWindow) Swap() {
	window.glfwWindow.SwapBuffers()
}
