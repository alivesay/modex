package gl

import (
	"errors"

	"github.com/alivesay/modex/core"
)

type Mesh struct {
	Data    []Vertex
	VBO     *VBO
	attribs []VertexAttrib
}

func NewMesh(usage VBOUsage, initialCapacity int) (*Mesh, error) {
	var data []Vertex = make([]Vertex, 0, initialCapacity)

	// TODO: handle shape specific attribs
	mesh := &Mesh{
		Data:    data,
		attribs: make([]VertexAttrib, 0, int(GetInstanceInfo().MaxVertexAttribs)),
	}

	var err error
	mesh.VBO, err = NewVBO(mesh.Data, mesh.attribs, usage)
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
	dataLen := len(mesh.Data)
	capLen := cap(mesh.Data)
	if dataLen == capLen {
		newData := make([]Vertex, dataLen, core.NP2(uint32(capLen+1)))
		copy(newData, mesh.Data)
		mesh.Data = newData
	}
	mesh.Data = append(mesh.Data, vertex)
}

func (mesh *Mesh) SyncBuffer() {
	mesh.VBO.UpdateBuffer(mesh.Data)
}

func (mesh *Mesh) ClearBuffer() {
	mesh.Data = mesh.Data[:0]
}
