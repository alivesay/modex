package gui

type WidgetContainerable interface {
	Parent() *Widget
	SetParent(*Widget)
	Children() []*Widget
	AddChildren(...*Widget)
	RemoveChildren(...*Widget)
}

type WidgetContainer struct {
	parent   *Widget
	children []*Widget
}

func (wc *WidgetContainer) Parent() *Widget {
	return wc.parent
}

func (wc *WidgetContainer) SetParent(parent *Widget) {
	wc.parent = parent
}

func (wc *WidgetContainer) Children() []*Widget {
	return wc.children
}

func (wc *WidgetContainer) AddChildren(...*Widget) {
	panic("unimplemented")
}

func (wc *WidgetContainer) RemoveChildren(...*Widget) {
	panic("unimplemented")
}
