package modex

import (
	"github.com/alivesay/modex/core"
	"github.com/alivesay/modex/gfx"
	"github.com/go-gl/glfw/v3.1/glfw"
)

type Modex struct {
	Running bool
	Video   *gfx.Video
}

func NewModex() *Modex {
	modex := &Modex{}

	if err := glfw.Init(); err != nil {
		core.Log(core.LOG_PANIC, err)
	}

	video, err := gfx.NewVideo()

	if err != nil {
		core.Log(core.LOG_PANIC, err)
		return nil
	}

	modex.Video = video

	return modex
}

func (modex *Modex) Destroy() {
	modex.Video.Destroy()
	glfw.Terminate()
}

func (modex *Modex) Boot() {
	modex.Running = true
}

func (modex *Modex) Shutdown() {
	modex.Running = false
}

func (modex *Modex) Update() {
}

func (modex *Modex) Render() {
	modex.Video.RenderBegin()
	modex.Video.RenderEnd()
}
