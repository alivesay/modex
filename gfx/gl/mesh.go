package gl

import (
	"errors"

	"github.com/alivesay/modex/core"
)

// Mesh manages structured vertex data.
type Mesh struct {
	vertices      []Vertex
	VBO           *VBO
	attribs       []VertexAttrib
	primitiveType GLPrimitiveType
}

var DefaultVertexAttribs = []VertexAttrib{
	VertexAttrib{0, 3, GLFloat, false, 5 * 4, 0},
	VertexAttrib{1, 2, GLFloat, false, 5 * 4, 3 * 4},
}

// NewMesh creates a new Mesh struct with a VBO capacity equal to initialCapacity.
func NewMesh(primitiveType GLPrimitiveType, usage VBOUsage, initialCapacity int) (*Mesh, error) {
	var vertices = make([]Vertex, 0, initialCapacity)

	// TODO: usage state
	mesh := &Mesh{
		primitiveType: primitiveType,
		vertices:      vertices,
		attribs:       DefaultVertexAttribs,
	}

	var err error
	mesh.VBO, err = NewVBO(mesh.vertices, mesh.primitiveType, mesh.attribs, usage)
	if err != nil {
		return nil, err
	}

	return mesh, nil
}

// Destroy implements a Mesh destructor.
func (mesh *Mesh) Destroy() {
	mesh.VBO.Destroy()
}

// TODO: why have the buffer and attribs both here and in VBO?

// AddVertexAttrib pushes a new VertexAttrib for use with this Mesh's VBO.
// Will replace DefaultVertices when called.
func (mesh *Mesh) AddVertexAttrib(attrib VertexAttrib) error {
	if &mesh.attribs == &DefaultVertexAttribs {
		mesh.attribs = make([]VertexAttrib, 0, int(GetInstanceInfo().MaxVertexAttribs))
	}

	if len(mesh.attribs) < int(GetInstanceInfo().MaxVertexAttribs) {
		mesh.attribs = append(mesh.attribs, attrib)
		if err := mesh.VBO.SetAttribs(mesh.attribs); err != nil {
			return err
		}
		return nil
	}

	return errors.New("GL_MAX_VERTEX_ATTRIBS exceeded")
}

// AddVertex appends current Vertex Data with new Vertex.
func (mesh *Mesh) AddVertex(vertex Vertex) {
	dataLen := len(mesh.vertices)
	capLen := cap(mesh.vertices)
	if dataLen == capLen {
		newData := make([]Vertex, dataLen, core.NP2(uint(capLen+1)))
		copy(newData, mesh.vertices)
		mesh.vertices = newData
	}
	mesh.vertices = append(mesh.vertices, vertex)
}

// SyncBuffer sets VBO buffer source to Mesh Data, possibly creating a new VBO.
func (mesh *Mesh) SyncBuffer() {
	mesh.VBO.UpdateBuffer(mesh.vertices)
}

// ClearBuffer removes all Mesh vertices.
func (mesh *Mesh) ClearBuffer() {
	mesh.vertices = mesh.vertices[:0]
}
