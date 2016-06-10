package modex

import (
	"github.com/alivesay/modex/core"
	"github.com/alivesay/modex/gfx"
	"github.com/go-gl/glfw/v3.1/glfw"
)

type Modex struct {
	Video *gfx.Video
	app   *core.Application
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

	modex.app = core.GetInstanceApplication()
	modex.Video = video

	return modex
}

func (modex *Modex) Destroy() {
	modex.Video.Destroy()
	glfw.Terminate()
}

func (modex *Modex) Boot() {
	modex.app.Running = true
	core.Log(core.LOG_NOTICE, "Booting...")
}

func (modex *Modex) Run() {
	core.Log(core.LOG_NOTICE, "Running...")
	for modex.app.ShutdownRequested == false {
		if modex.app.Running {
			modex.update()
			modex.render()
		}
	}
}

func (modex *Modex) Pause() {
	modex.app.Running = false
}

func (modex *Modex) Shutdown() {
	modex.app.Running = false
	core.Log(core.LOG_NOTICE, "Shutting down...")
}

func (modex *Modex) update() {
	// initiate update cascade
}

func (modex *Modex) render() {
	modex.Video.Render()
}
