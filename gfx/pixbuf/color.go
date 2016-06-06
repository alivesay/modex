package pixbuf

import (
	"image/color"
)

type RGBA32 struct {
	color.Color
	Packed uint32
}

const (
	RMASK   = 0xFF000000
	GMASK   = 0x00FF0000
	BMASK   = 0x0000FF00
	AMASK   = 0x000000FF
	RGBMASK = 0x00FFFFFF
	RSHIFT  = 24
	GSHIFT  = 16
	BSHIFT  = 8
	ASHIFT  = 0
)

func (rgba32 *RGBA32) RGBA() (r, g, b, a uint32) {
	return rgba32.Packed >> RSHIFT & 0xFF,
		rgba32.Packed >> GSHIFT & 0xFF,
		rgba32.Packed >> BSHIFT & 0xFF,
		rgba32.Packed >> ASHIFT & 0xFF
}

func (rgba32 *RGBA32) FromRGBA8(r, g, b, a uint8) {
	rgba32.Packed = (uint32(r) << RSHIFT) + (uint32(g) << GSHIFT) + (uint32(b) << BSHIFT) + (uint32(a) << ASHIFT)
}

func (rgba32 *RGBA32) FromRGBA(r, g, b uint8) {
	rgba32.FromRGBA8(r, g, b, 0xFF)
}

func (rgba32 *RGBA32) R() uint8 {
	return uint8((rgba32.Packed & RMASK) >> RSHIFT)
}

func (rgba32 *RGBA32) G() uint8 {
	return uint8((rgba32.Packed & GMASK) >> GSHIFT)
}

func (rgba32 *RGBA32) B() uint8 {
	return uint8((rgba32.Packed & BMASK) >> BSHIFT)
}

func (rgba32 *RGBA32) A() uint8 {
	return uint8((rgba32.Packed & AMASK) >> ASHIFT)
}

func (rgba32 *RGBA32) Rn() float32 {
	return float32((rgba32.Packed&RMASK)>>RSHIFT) / 255.0
}

func (rgba32 *RGBA32) Gn() float32 {
	return float32((rgba32.Packed&GMASK)>>GSHIFT) / 255.0
}

func (rgba32 *RGBA32) Bn() float32 {
	return float32((rgba32.Packed&BMASK)>>BSHIFT) / 255.0
}

func (rgba32 *RGBA32) An() float32 {
	return float32((rgba32.Packed&AMASK)>>ASHIFT) / 255.0
}
