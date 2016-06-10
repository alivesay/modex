package gl

import (
	"bytes"
	gl "github.com/go-gl/glow/gl"
)

type VertexAttrib struct {
	Index      uint32
	Size       int32
	Type       GLType
	Normalized bool
	Stride     uint32
	Offset     uintptr
}

type VBOUsage int

const (
	StaticDraw  VBOUsage = gl.STATIC_DRAW
	DynamicDraw VBOUsage = gl.DYNAMIC_DRAW
	StreamDraw  VBOUsage = gl.STREAM_DRAW
)

type VBO struct {
	glVBOIDs      [2]uint32
	bufferUsage   VBOUsage
	bufferSize    uint32
	bufferAttribs []VertexAttrib
	// TODO: support vbo doublebuffering
	//currentBuffer uint32
	buffer  *bytes.Buffer
	attribs []VertexAttrib
}

func NewVBO(dataSize uint32, data []Vertex, attribs []VertexAttrib, vboUsage VBOUsage) *VBO {
	return &VBO{}
}

func (vbo *VBO) Destroy() {
}
