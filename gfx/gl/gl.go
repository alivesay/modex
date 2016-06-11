package gl

import (
	"github.com/alivesay/modex/core"
	gl "github.com/go-gl/glow/gl"
)

type GLType int32

const (
	GLByte         GLType = gl.BYTE
	GLUnsignedByte GLType = gl.UNSIGNED_BYTE
	GLShort        GLType = gl.SHORT
	GLInt          GLType = gl.INT
	GLUnsignedInt  GLType = gl.UNSIGNED_INT
	GLFloat        GLType = gl.FLOAT
	GLFixed        GLType = gl.FIXED
)

type State struct {
	GLInitialized bool
	Info          *Info
}

func (state *State) Init() {
	if state.GLInitialized {
		core.Log(core.LogErr, "GL already initialized")
		return
	}

	if err := gl.Init(); err != nil {
		core.Log(core.LogPanic, err)
	}

	state.Info = GetInstanceInfo()
	core.Log(core.LogDebug, state.Info)

	state.GLInitialized = true
}

func (state *State) Destroy() {
}
