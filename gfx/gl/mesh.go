package gl

import "errors"

// TODO: best initial value?
const initialVertexCount uint32 = 64
const maxVertexCount uint32 = 65536

type Mesh struct {
	Data    []Vertex
	VBO     *VBO
	attribs []VertexAttrib
}

func NewMesh(size uint32) *Mesh {
	var data []Vertex = make([]Vertex, initialVertexCount, maxVertexCount)

	data = append(data, Vertex{[3]float32{160.0, 10.0, 0.0}, [4]float32{0.0, 0.0, 0.0, 0.0}})
	data = append(data, Vertex{[3]float32{310.0, 230.0, 0.0}, [4]float32{0.0, 0.0, 0.0, 0.0}})
	data = append(data, Vertex{[3]float32{10.0, 230.0, 0.0}, [4]float32{0.0, 0.0, 0.0, 0.0}})

	mesh := &Mesh{
		Data: data,
		attribs: []VertexAttrib{
			{0, 3, GLFloat, false, 3, 0},
			{1, 4, GLFloat, false, 4, 3},
		},
	}

	mesh.VBO = NewVBO(uint32(len(mesh.Data)), mesh.Data, mesh.attribs, DynamicDraw)

	return mesh
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
