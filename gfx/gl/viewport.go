package gl

import (
	"errors"

	"github.com/alivesay/modex/gfx/gui"
	"github.com/alivesay/modex/gfx/pixbuf"
	gogl "github.com/go-gl/gl/all-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

// A Viewport manages OpenGL viewports within the scene.
type Viewport struct {
	ProjMat     *mgl32.Mat4
	RenderState *RenderState
	BgColor     *pixbuf.RGBA32
	gui.Rectangle
	Renderable
}

// NewViewport constructs a new managed Viewport.
func NewViewport(x, y, width, height int32) *Viewport {
	viewport := &Viewport{
		ProjMat:     &mgl32.Mat4{},
		BgColor:     &pixbuf.RGBA32{Packed: 0x3366CCFF},
		RenderState: NewRenderState(),
		Rectangle:   gui.NewRectangle(x, y, width, height),
	}

	viewport.RenderState.AddCapability(gogl.CULL_FACE)
	viewport.RenderState.AddCapability(gogl.BLEND)

	return viewport
}

// SetPerspective establishes a default projection matrix.
func (viewport *Viewport) SetPerspective() error {
	if viewport.Width() <= 0 || viewport.Height() <= 0 {
		return errors.New("viewport dimensions must be positive values")
	}

	gogl.Viewport(viewport.X(), viewport.Y(), viewport.Width(), viewport.Height())
	projMat := mgl32.Perspective(mgl32.DegToRad(60.0), float32(viewport.Width())/float32(viewport.Height()), 0.1, 2000.0)
	viewport.ProjMat = &projMat

	return nil
}

// SetOrtho2D establishes an orthographic projection matrix.
func (viewport *Viewport) SetOrtho2D() error {
	maxDims := GetInstanceInfo().MaxViewportDims
	if viewport.Width()-viewport.X() <= maxDims[0] && viewport.Height()-viewport.Y() <= maxDims[1] {
		gogl.Viewport(viewport.X(), viewport.Y(), viewport.Width(), viewport.Height())
		projMat := mgl32.Ortho2D(float32(viewport.X()), float32(viewport.Width()), float32(viewport.Height()), float32(viewport.Y()))
		viewport.ProjMat = &projMat

		return nil
	}

	return errors.New("GL_MAX_VIEWPORT_DIMS exceeded")
}

// TODO:

// Render needs to move somewhere else.
func (viewport *Viewport) Render() {
	gogl.Enable(gogl.CULL_FACE)
	//	gogl.Enable(gogl.DEPTH_TEST)
	gogl.Enable(gogl.BLEND)
	gogl.BlendFunc(gogl.SRC_ALPHA, gogl.ONE_MINUS_SRC_ALPHA)
	gogl.CullFace(gogl.BACK)
	viewport.Clear()
	// Render members
	//	gogl.Disable(gogl.DEPTH_TEST)
	//	gogl.Disable(gogl.CULL_FACE)
}

// TODO:

// Viewport clears the screen.. should it clear just the viewport? or move elsewhere.
func (viewport *Viewport) Clear() {
	gogl.ClearColor(viewport.BgColor.Rn(), viewport.BgColor.Gn(), viewport.BgColor.Bn(), viewport.BgColor.An())
	gogl.Clear(gogl.COLOR_BUFFER_BIT | gogl.DEPTH_BUFFER_BIT)
}
