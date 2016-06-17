package gui

type Position [3]int
type Dimensions [2]int

type Rectangleable interface {
	X() int
	Y() int
	Z() int
	SetX() int
	SetY() int
	SetZ() int
	Position() Position
	SetPosition(Position)
	Dimensions() Dimensions
	SetDimensions(Dimensions)
	Width() int
	Height() int
	SetWidth(int)
	SetHeight(int)
	ContainsPoint(int)
}

type Rectangle struct {
	Rectangleable
	position   Position
	dimensions Dimensions
}

func (rect *Rectangle) X() int {
	return rect.position[0]
}

func (rect *Rectangle) Y() int {
	return rect.position[1]
}

func (rect *Rectangle) Z() int {
	return rect.position[2]
}

func (rect *Rectangle) SetX(x int) {
	rect.position[0] = x
}

func (rect *Rectangle) SetY(y int) {
	rect.position[1] = y
}

func (rect *Rectangle) SetZ(z int) {
	rect.position[2] = z
}

func (rect *Rectangle) Position() Position {
	return rect.position
}

func (rect *Rectangle) SetPosition(position Position) {
	rect.position = position
}

func (rect *Rectangle) Dimensions() Dimensions {
	return rect.dimensions
}

func (rect *Rectangle) SetDimensions(dimensions Dimensions) {
	rect.dimensions = dimensions
}

func (rect *Rectangle) Width() int {
	return rect.dimensions[0]
}

func (rect *Rectangle) Height() int {
	return rect.dimensions[1]
}

func (rect *Rectangle) SetWidth(width int) {
	rect.dimensions[0] = width
}

func (rect *Rectangle) SetHeight(height int) {
	rect.dimensions[1] = height
}
