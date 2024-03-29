package gl

import (
	gogl "github.com/go-gl/gl/all-core/gl"
	"sort"
)

type RenderState struct {
	Capabilities []int
	Shader       *Shader
	VBO          *VBO
}

func NewRenderState() *RenderState {
	return &RenderState{
		Capabilities: make([]int, 0),
		Shader:       nil,
		VBO:          nil,
	}
}

func (state *RenderState) Destroy() {
}

func (state *RenderState) AddCapability(capability int) {
	state.Capabilities = append(state.Capabilities, capability)
	sort.Ints(state.Capabilities)
}

func (state *RenderState) RemoveCapability(capability int) {
	i := sort.SearchInts(state.Capabilities, capability)
	if i < len(state.Capabilities) && state.Capabilities[i] == capability {
		state.Capabilities = append(state.Capabilities[:i], state.Capabilities[i+1:]...)
	}
}

func (state *RenderState) Enable() {
	// enable caps
	// bind shader
	// bind texture
	// bind VBO

	for cap := range state.Capabilities {
		gogl.Enable(uint32(cap))
	}
}

func (state *RenderState) Disable() {
	for cap := range state.Capabilities {
		gogl.Disable(uint32(cap))
	}
}
