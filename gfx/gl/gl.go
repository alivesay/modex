package gl

import (
	"github.com/alivesay/modex/core"
	gogl "github.com/go-gl/gl/all-core/gl"
)

type GLType uint32

const (
	GLByte          GLType = gogl.BYTE
	GLUnsignedByte         = gogl.UNSIGNED_BYTE
	GLShort                = gogl.SHORT
	GLInt                  = gogl.INT
	GLUnsignedShort        = gogl.UNSIGNED_SHORT
	GLUnsignedInt          = gogl.UNSIGNED_INT
	GLFloat                = gogl.FLOAT
	GLFixed                = gogl.FIXED
)

// GLPrimitiveType represents OpenGL primitive types
type GLPrimitiveType int32

// Wrappers for supported Opengl primitive types.
const (
	GLPoints        GLPrimitiveType = gogl.POINTS
	GLLines                         = gogl.LINES
	GLLineStrip                     = gogl.LINE_STRIP
	GLLineLoop                      = gogl.LINE_LOOP
	GLTriangles                     = gogl.TRIANGLES
	GLTriangleStrip                 = gogl.TRIANGLE_STRIP
	GLTriangleFan                   = gogl.TRIANGLE_FAN
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
