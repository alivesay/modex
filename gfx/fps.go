package gfx

import (
	"fmt"

	"github.com/go-gl/glfw/v3.1/glfw"
)

type FPS struct {
	lastTime float64
	nbFrames int
}

func NewFPS() *FPS {
	fps := &FPS{
		lastTime: glfw.GetTime(),
		nbFrames: 0,
	}

	return fps
}

// TODO:

func (fps *FPS) Update(window *glfw.Window) {
	currentTime := glfw.GetTime()
	fps.nbFrames++

	if currentTime-fps.lastTime >= 1.0 {
		window.SetTitle(fmt.Sprintf("%.3f ms/frame, %d fps", 1000.0/float64(fps.nbFrames), fps.nbFrames))
		fps.nbFrames = 0
		fps.lastTime += 1.0
	}
}
