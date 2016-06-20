package gl

import (
	"github.com/alivesay/modex/core"
	gogl "github.com/go-gl/gl/all-core/gl"
)

type GLType int32

const (
	GLByte          GLType = gogl.BYTE
	GLUnsignedByte  GLType = gogl.UNSIGNED_BYTE
	GLShort         GLType = gogl.SHORT
	GLInt           GLType = gogl.INT
	GLUnsignedShort GLType = gogl.UNSIGNED_SHORT
	GLUnsignedInt   GLType = gogl.UNSIGNED_INT
	GLFloat         GLType = gogl.FLOAT
	GLFixed         GLType = gogl.FIXED
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

	if err := gogl.Init(); err != nil {
		core.Log(core.LogPanic, err)
	}

	state.Info = GetInstanceInfo()
	core.Log(core.LogDebug, state.Info)

	state.GLInitialized = true
}

func (state *State) Destroy() {
}
