package events

import (
	"github.com/go-gl/glfw/v3.1/glfw"
)

func Poll() {
	glfw.PollEvents()
}
