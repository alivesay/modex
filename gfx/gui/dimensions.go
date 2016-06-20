package gui

type Dimensionsable interface {
	Width() int32
	Height() int32
	SetWidth(int32)
	SetHeight(int32)
}

type Dimensions struct {
	width  int32
	height int32
}

func NewDimensions(width, height int32) Dimensions {
	return Dimensions{width: width, height: height}
}

func (dims Dimensions) Width() int32 {
	return dims.width
}

func (dims Dimensions) Height() int32 {
	return dims.height
}

func (dims Dimensions) SetWidth(width int32) {
	dims.width = width
}

func (dims Dimensions) SetHeight(height int32) {
	dims.height = height
}
