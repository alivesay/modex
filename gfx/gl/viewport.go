package gl

// TODO: use prim.Dimensions instead of rect here.

import (
	"errors"
	"image"

	"github.com/alivesay/modex/gfx/pixbuf"

	gogl "github.com/go-gl/gl/all-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

// A Viewport manages OpenGL viewports within the scene.
type Viewport struct {
	ProjMat     *mgl32.Mat4
	RenderState *RenderState
	BgColor     *pixbuf.RGBA32
	Rect        image.Rectangle
	Renderable
}

// NewViewport constructs a new managed Viewport.
func NewViewport(x, y, width, height int) *Viewport {
	v := &Viewport{
		ProjMat:     &mgl32.Mat4{},
		BgColor:     &pixbuf.RGBA32{Packed: 0x3366CCFF},
		RenderState: NewRenderState(),
		Rect:        image.Rect(x, y, width, height),
	}

	v.RenderState.AddCapability(gogl.CULL_FACE)
	v.RenderState.AddCapability(gogl.BLEND)

	return v
}

// SetPerspective establishes a default projection matrix.
func (v *Viewport) SetPerspective() error {
	if v.Rect.Dx() <= 0 || v.Rect.Dy() <= 0 {
		return errors.New("viewport dimensions must be positive values")
	}

	gogl.Viewport(int32(v.Rect.Min.X), int32(v.Rect.Min.Y), int32(v.Rect.Max.X), int32(v.Rect.Max.Y))
	projMat := mgl32.Perspective(mgl32.DegToRad(60.0), float32(v.Rect.Dx())/float32(v.Rect.Dy()), 0.1, 2000.0)
	v.ProjMat = &projMat

	return nil
}

// SetOrtho2D establishes an orthographic projection matrix.
func (v *Viewport) SetOrtho2D() error {
	maxDims := GetInstanceInfo().MaxViewportDims
	if v.Rect.Dx() <= int(maxDims[0]) && v.Rect.Dy() <= int(maxDims[1]) {
		gogl.Viewport(int32(v.Rect.Min.X), int32(v.Rect.Min.Y), int32(v.Rect.Max.X), int32(v.Rect.Max.Y))
		projMat := mgl32.Ortho2D(float32(v.Rect.Min.X), float32(v.Rect.Max.X), float32(v.Rect.Max.Y), float32(v.Rect.Min.Y))
		v.ProjMat = &projMat

		return nil
	}

	return errors.New("GL_MAX_VIEWPORT_DIMS exceeded")
}

// TODO:

// Render needs to move somewhere else.
func (v *Viewport) Render() {
	gogl.Enable(gogl.CULL_FACE)
	//	gogl.Enable(gogl.DEPTH_TEST)
	gogl.BlendFunc(gogl.SRC_ALPHA, gogl.ONE_MINUS_SRC_ALPHA)
	gogl.Enable(gogl.BLEND)
	gogl.CullFace(gogl.BACK)
	v.Clear()
	// Render members
	//	gogl.Disable(gogl.DEPTH_TEST)
	//	gogl.Disable(gogl.CULL_FACE)
}

// TODO:

// Viewport clears the screen.. should it clear just the viewport? or move elsewhere.
func (v *Viewport) Clear() {
	gogl.ClearColor(v.BgColor.Rn(), v.BgColor.Gn(), v.BgColor.Bn(), v.BgColor.An())
	gogl.Clear(gogl.COLOR_BUFFER_BIT | gogl.DEPTH_BUFFER_BIT)
}
