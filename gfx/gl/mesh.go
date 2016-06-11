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

func NewMesh() (*Mesh, error) {
	var data []Vertex = make([]Vertex, 0, maxVertexCount)

	// TODO: handle shape specific attribs
	mesh := &Mesh{
		Data:    data,
		attribs: make([]VertexAttrib, 0, int(GetInstanceInfo().MaxVertexAttribs)),
	}

	var err error
	mesh.VBO, err = NewVBO(mesh.Data, mesh.attribs, DynamicDraw)
	if err != nil {
		return nil, err
	}

	return mesh, nil
}

func (mesh *Mesh) Destroy() {
	mesh.VBO.Destroy()
}

func (mesh *Mesh) AddVertexAttrib(attrib VertexAttrib) error {
	if len(mesh.attribs) < int(GetInstanceInfo().MaxVertexAttribs) {
		mesh.attribs = append(mesh.attribs, attrib)
		if err := mesh.VBO.UpdateAttribs(mesh.attribs); err != nil {
			return err
		}
		return nil
	}

	return errors.New("GL_MAX_VERTEX_ATTRIBS exceeded")
}

func (mesh *Mesh) AddVertex(vertex Vertex) {
	// TODO: check max vertexes
	mesh.Data = append(mesh.Data, vertex)
	mesh.VBO.UpdateBuffer(mesh.Data)
}
