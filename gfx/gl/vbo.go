package gl

import (
	"errors"
	gl "github.com/go-gl/glow/gl"
)

type VertexAttrib struct {
	Index      uint32
	Size       int32
	Type       GLType
	Normalized bool
	Stride     uint32
	Offset     int
}

type VBOUsage uint32

const (
	StaticDraw  VBOUsage = gl.STATIC_DRAW
	DynamicDraw VBOUsage = gl.DYNAMIC_DRAW
	StreamDraw  VBOUsage = gl.STREAM_DRAW
)

type VBO struct {
	glVBOID       uint32
	glVAOID       uint32
	bufferUsage   VBOUsage
	bufferSize    int
	bufferAttribs []VertexAttrib
	// TODO: support vbo doublebuffering
	//currentBuffer uint32
	buffer  []Vertex
	attribs []VertexAttrib
}

func NewVBO(dataSize uint32, data []Vertex, attribs []VertexAttrib, vboUsage VBOUsage) (*VBO, error) {
	vbo := &VBO{
		bufferUsage: vboUsage,
		bufferSize:  int(dataSize),
		attribs:     attribs,
		buffer:      data,
	}

	gl.GenBuffers(1, &vbo.glVBOID)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo.glVBOID)

	if gl.IsBuffer(vbo.glVBOID) == false {
		return nil, errors.New("failed to bind buffer")
	}
	for _, attrib := range vbo.attribs {
		gl.EnableVertexAttribArray(attrib.Index)
		gl.VertexAttribPointer(attrib.Index, attrib.Size, uint32(attrib.Type), attrib.Normalized, int32(attrib.Stride), gl.PtrOffset(attrib.Offset))
	}

	gl.BufferData(gl.ARRAY_BUFFER, len(vbo.buffer)*24, nil, uint32(vbo.bufferUsage))
	gl.BufferSubData(gl.ARRAY_BUFFER, 0, 3*24, gl.Ptr(vbo.buffer))

	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	return vbo, nil
}

func (vbo *VBO) Destroy() {
	gl.DeleteBuffers(1, &vbo.glVBOID)
}

func (vbo *VBO) Bind() {
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo.glVBOID)
	gl.BufferData(gl.ARRAY_BUFFER, 3*24, nil, uint32(vbo.bufferUsage))
	gl.BufferSubData(gl.ARRAY_BUFFER, 0, 3*24, gl.Ptr(vbo.buffer))

	for _, attrib := range vbo.attribs {
		gl.EnableVertexAttribArray(attrib.Index)
		gl.VertexAttribPointer(attrib.Index, attrib.Size, uint32(attrib.Type), attrib.Normalized, int32(attrib.Stride), gl.PtrOffset(attrib.Offset))
	}

}

func (vbo *VBO) Unbind() {
	for _, attrib := range vbo.attribs {
		gl.DisableVertexAttribArray(attrib.Index)
	}
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
}

func (vbo *VBO) Render() {
	gl.DrawArrays(gl.TRIANGLES, 0, 3)
}
