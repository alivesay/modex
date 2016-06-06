package gl

import (
	"github.com/alivesay/modex/core"
	gl "github.com/go-gl/glow/gl"
)

type State struct {
	GLInitialized bool
	Info          *Info
}

func (state *State) Init() {
	if state.GLInitialized {
		core.Log(core.LOG_ERR, "GL already initialized")
		return
	}

	if err := gl.Init(); err != nil {
		core.Log(core.LOG_PANIC, err)
	}

	state.Info = NewInfo()
	core.Log(core.LOG_DEBUG, state.Info)

	state.GLInitialized = true
}

func (state *State) Destroy() {
}
