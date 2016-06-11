package gl

import (
	"errors"
	"github.com/alivesay/modex/gfx/pixbuf"
	gles2 "github.com/go-gl/gl/v3.1/gles2"
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
	maxDims := GetInstanceInfo().MaxViewportDims
	if viewport.W-viewport.X <= maxDims[0] && viewport.H-viewport.Y <= maxDims[1] {
		gles2.Viewport(viewport.X, viewport.Y, viewport.W, viewport.H)
		projMat := mgl32.Ortho2D(float32(viewport.X), float32(viewport.W), float32(viewport.H), float32(viewport.Y))
		viewport.ProjMat = &projMat

		return nil
	}

	return errors.New("GL_MAX_VIEWPORT_DIMS exceeded")
}

func (viewport *Viewport) Clear(color *pixbuf.RGBA32) {
	gles2.ClearColor(color.Rn(), color.Gn(), color.Bn(), color.An())
	gles2.Clear(gles2.COLOR_BUFFER_BIT | gles2.DEPTH_BUFFER_BIT)
}
