package gui

type Positionable interface {
	X() int32
	Y() int32
	SetX(int32)
	SetY(int32)
}

type Position struct {
	x int32
	y int32
}

func NewPosition(x, y int32) Position {
	return Position{x: x, y: y}
}

func (pos Position) X() int32 {
	return pos.x
}

func (pos Position) Y() int32 {
	return pos.y
}

func (pos Position) SetX(x int32) {
	pos.x = x
}

func (pos Position) SetY(y int32) {
	pos.y = y
}
