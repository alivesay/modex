package gl

import (
	"errors"
	"github.com/alivesay/modex/gfx/pixbuf"
	gles2 "github.com/go-gl/gl/v3.1/gles2"
	"github.com/go-gl/mathgl/mgl32"
)

type Viewport struct {
	X           int32
	Y           int32
	W           int32
	H           int32
	ProjMat     *mgl32.Mat4
	BgColor     *pixbuf.RGBA32
	RenderState *RenderState
	Renderable
}

func NewViewport(x, y, w, h int32) *Viewport {
	viewport := &Viewport{
		X:           x,
		Y:           y,
		W:           w,
		H:           h,
		ProjMat:     &mgl32.Mat4{},
		BgColor:     &pixbuf.RGBA32{Packed: 0x3366CCFF},
		RenderState: NewRenderState(),
	}

	viewport.RenderState.AddCapability(gles2.CULL_FACE)

	return viewport
}

func (viewport *Viewport) Destroy() {
}

func (viewport *Viewport) SetPerspective() error {
	if viewport.W <= 0 || viewport.H <= 0 {
		return errors.New("viewport dimensions must be positive values")
	}

	gles2.Viewport(viewport.X, viewport.Y, viewport.W, viewport.H)
	projMat := mgl32.Perspective(mgl32.DegToRad(60.0), float32(viewport.W)/float32(viewport.H), 0.1, 2000.0)
	viewport.ProjMat = &projMat

	return nil
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

func (viewport *Viewport) Clear() {
	gles2.ClearColor(viewport.BgColor.Rn(), viewport.BgColor.Gn(), viewport.BgColor.Bn(), viewport.BgColor.An())
	gles2.Clear(gles2.COLOR_BUFFER_BIT | gles2.DEPTH_BUFFER_BIT)
}

func (viewport *Viewport) Render() {
	//	gles2.Enable(gles2.CULL_FACE)
	//	gles2.Enable(gles2.DEPTH_TEST)
	gles2.Enable(gles2.BLEND)
	gles2.BlendFunc(gles2.SRC_ALPHA, gles2.ONE_MINUS_SRC_ALPHA)
	viewport.Clear()
	// Render members
	//	gles2.Disable(gles2.DEPTH_TEST)
	//	gles2.Disable(gles2.CULL_FACE)
}
