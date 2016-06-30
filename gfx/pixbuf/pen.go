package pixbuf

type Pen struct {
	Color RGBA32
	Width int
}

var WhitePen = NewPen(RGBA32{Packed: 0xFFFFFFFF})
var BlackPen = NewPen(RGBA32{Packed: 0x000000FF})
var ClearPen = NewPen(RGBA32{Packed: 0x00000000})

func NewPen(color RGBA32) Pen {
	return Pen{Color: color}
}
