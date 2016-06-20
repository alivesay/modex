package gui

type Rectangleable interface {
	Positionable
	Dimensionsable
	Position() Position
	SetPosition(Position)
	Dimensions() Dimensions
	SetDimensions(Dimensions)
	ContainsPoint(int32)
}

type Rectangle struct {
	Position
	Dimensions
}

func NewRectangle(x, y, width, height int32) Rectangle {
	return Rectangle{
		Position:   Position{x: x, y: y},
		Dimensions: Dimensions{width: width, height: height},
	}
}

func (rect Rectangle) ContainsPoint(x, y int32) bool {
	panic("not implemented")
}
