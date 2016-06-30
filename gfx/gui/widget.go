package gui

import (
	"image"

	"github.com/alivesay/modex/gfx/prim"
)

type Drawable interface {
	Draw()
}

type Widgetable interface {
	prim.Positionable
	prim.Dimensionable
	Containerable
	Drawable

	Name() string
	SetName(string)

	Visible() bool
	SetVisible(bool)

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
	Dimensions
	WidgetContainer

	Surface *image.NRGBA

	name      string
	visible   bool
	autoScale bool
}

func NewWidget() *Widget {
	return &Widget{
		dims:              prim.Dimensions,
		autoScaleToParent: true,
	}
}

func (widget *Widget) Name() string {
	return widget.name
}

func (widget *Widgte) SetName(name string) {
	widget.name = name
}

func (widget *Widget) Visible() bool {
	return widget.visible
}

func (widget *Widget) SetVisible(visible bool) {
	widget.visible = visible
}

func (widget *Widget) AutoScaleToParent() bool {
	return widget.autoScaleToParent
}

func (widget *Widget) SetVisible(autoScale bool) {
	widget.autoScaleToParent = autoScale
}
