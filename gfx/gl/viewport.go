package gl

import (
	"errors"
	"fmt"
	"github.com/alivesay/modex/gfx/pixbuf"
	gl "github.com/go-gl/glow/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type Viewport struct {
	X       int32
	Y       int32
	W       int32
	H       int32
	ProjMat *mgl32.Mat4
}

func (viewport *Viewport) SetOrtho2D() error {
	fmt.Println(GetInstanceInfo())
	maxDims := GetInstanceInfo().MaxViewportDims
	if viewport.W-viewport.X <= maxDims[0] && viewport.H-viewport.Y <= maxDims[1] {
		gl.Viewport(viewport.X, viewport.Y, viewport.W, viewport.H)

		projMat := mgl32.Ortho2D(float32(viewport.X), float32(viewport.W), float32(viewport.Y), float32(viewport.H))
		viewport.ProjMat = &projMat

		return nil
	}

	return errors.New("GL_MAX_VIEWPORT_DIMS exceeded")
}

func (viewport *Viewport) Clear(color *pixbuf.RGBA32) {
	gl.ClearColor(color.Rn(), color.Gn(), color.Bn(), color.An())
	gl.Clear(gl.COLOR_BUFFER_BIT)
}
