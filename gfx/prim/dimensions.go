package prim

type Dimensionable interface {
	Width() int
	Height() int
	SetWidth(int)
	SetHeight(int)
}

type Dimensions struct {
	height, width int
}

func (dims *Dimensions) Width() int {
	return dims.width
}

func (dims *Dimensions) Height() int {
	return dims.height
}

func (dims *Dimensions) SetWidth(w int) {
	dims.width = w
}

func (dims *Dimensions) SetHeight(h int) {
	dims.height = h
}
