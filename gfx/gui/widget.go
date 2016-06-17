package gui

type Widgetable interface {
	Rectangleable

	Parent() *Widget
	SetParent(*Widget)
	Children() []*Widget
	AddChildren(...*Widget)
	RemoveChildren(...*Widget)

	ClientX() int
	ClientY() int
	ClientZ() int
	ScreenX() int
	ScreenY() int

	MinWidth() int
	MinHeight() int
	SetMinWidth(int)
	SetMinHeight(int)

	BorderWidth() int
	SetBorderWidth(int)

	AutoScaleToParent() bool
	SetAutoScaleToParent(bool)
}

type Widget struct {
	Rectangle
}
