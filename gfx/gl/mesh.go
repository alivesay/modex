package gl

import (
	"errors"
)

// TODO: best initial value?
const initialVertexCount uint32 = 64
const maxVertexCount uint32 = 65536

type Mesh struct {
	Data    []Vertex
	VBO     *VBO
	attribs []VertexAttrib
}

func NewMesh(size uint32) (*Mesh, error) {
	var data []Vertex = make([]Vertex, 0, maxVertexCount)

	data = append(data, Vertex{160.0, 0.0, 0.0})
	data = append(data, Vertex{320.0, 240.0, 0.0})
	data = append(data, Vertex{0.0, 240, 0.0})

	mesh := &Mesh{
		Data: data,
		attribs: []VertexAttrib{
			{0, 3, GLFloat, false, 0, 0},
			//			{1, 4, GLFloat, false, 32, 4},
		},
	}

	var err error
	mesh.VBO, err = NewVBO(uint32(len(mesh.Data)), mesh.Data, mesh.attribs, DynamicDraw)
	if err != nil {
		return nil, err
	}

	return mesh, nil
}

func (mesh *Mesh) addVertexAttrib(attrib VertexAttrib) error {
	if len(mesh.attribs) < int(GetInstanceInfo().MaxVertexAttribs) {
		mesh.attribs = append(mesh.attribs, attrib)
		return nil
	}

	return errors.New("GL_MAX_VERTEX_ATTRIBS exceeded")
}

func Destroy() {
}
