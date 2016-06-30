package pixbuf

type BrushStyle int

const (
	SolidBrush BrushStyle = iota
)

type Brush struct {
	Color RGBA32
	Style BrushStyle
}

var WhiteSolidBrush = NewBrush(RGBA32{Packed: 0xFFFFFFFF}, SolidBrush)
var BlackSolidBrush = NewBrush(RGBA32{Packed: 0x000000FF}, SolidBrush)
var ClearSolidBrush = NewBrush(RGBA32{Packed: 0x00000000}, SolidBrush)

func NewBrush(color RGBA32, style BrushStyle) Brush {
	return Brush{Color: color, Style: style}
}
