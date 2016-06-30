package prim

type Positionable interface {
	X() int
	Y() int
	SetX(int)
	SetY(int)
}

type Position struct {
	x, y int
}

func (pos *Position) X() int {
	return pos.x
}

func (pos *Position) Y() int {
	return pos.y
}

func (pos *Position) SetX(x int) {
	pos.x = x
}

func (pos *Position) SetY(y int) {
	pos.y = y
}
