package gl

import (
	"github.com/alivesay/modex/core"
	gles2 "github.com/go-gl/gl/v3.1/gles2"
)

type GLType int32

const (
	GLByte          GLType = gles2.BYTE
	GLUnsignedByte  GLType = gles2.UNSIGNED_BYTE
	GLShort         GLType = gles2.SHORT
	GLInt           GLType = gles2.INT
	GLUnsignedShort GLType = gles2.UNSIGNED_SHORT
	GLUnsignedInt   GLType = gles2.UNSIGNED_INT
	GLFloat         GLType = gles2.FLOAT
	GLFixed         GLType = gles2.FIXED
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

	if err := gles2.Init(); err != nil {
		core.Log(core.LogPanic, err)
	}

	state.Info = GetInstanceInfo()
	core.Log(core.LogDebug, state.Info)

	state.GLInitialized = true
}

func (state *State) Destroy() {
}
