package gl

type Renderable interface {
	Render()
	RenderState() *RenderState
	SetRenderState(*RenderState)
}

// TODO: need a default renderer that just draws a rect

type Renderer struct {
	// rendertree
}

func NewRenderer() *Renderer {
	return &Renderer{}
}

func (r *Renderer) Render() {
}
