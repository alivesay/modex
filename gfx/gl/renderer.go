package gl

import (
	"github.com/alivesay/modex/gfx/pixbuf"
	gl "github.com/go-gl/glow/gl"
)

type Renderer struct {
	BgColor *pixbuf.RGBA32
	// rendertree
}

func NewRenderer() *Renderer {
	return &Renderer{BgColor: &pixbuf.RGBA32{Packed: 0x3366CCFF}}
}

func (r *Renderer) RenderBegin() {
	gl.ClearColor(r.BgColor.Rn(), r.BgColor.Gn(), r.BgColor.Bn(), r.BgColor.An())
	gl.Clear(gl.COLOR_BUFFER_BIT)
}

func (r *Renderer) RenderEnd() {
	// Check for GL errors
}
