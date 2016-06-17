package gl

type SceneObjectable interface {
	Renderable
	Parent() *SceneObject
	SetParent(*SceneObject)
	Children() []*SceneObject
	AddChildren(children ...SceneObject)
	RemoveChildren(children ...SceneObject)
	Name() string
	SetName(string)
}

type SceneObject struct {
	SceneObjectable
	Parent   *SceneObject
	Children []SceneObject
}

type Scene struct {
}
