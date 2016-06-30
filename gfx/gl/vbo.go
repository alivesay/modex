package gl

import (
	"errors"

	"github.com/alivesay/modex/core"

	gogl "github.com/go-gl/gl/all-core/gl"
)

// TODO: should i not use GLType here?

// VertexAttrib holds values matching GLVertexAttribPointer params.
type VertexAttrib struct {
	Index      uint32
	Size       int32
	Type       GLType
	Normalized bool
	Stride     int32
	Offset     int
}

// VBOUsage represents OpenGL VBO usage types.
type VBOUsage uint32

// VBOUsage constants.
const (
	StaticDraw  VBOUsage = gogl.STATIC_DRAW
	DynamicDraw          = gogl.DYNAMIC_DRAW
	StreamDraw           = gogl.STREAM_DRAW
)

// VBO maintains state for a managed OpenGL Vertex Buffer Object.
type VBO struct {
	glVBOID        uint32
	bufferUsage    VBOUsage
	bufferCapacity int
	bufferAttribs  []VertexAttrib
	// TODO: support vbo doublebuffering
	//currentBuffer uint32
	buffer        []Vertex
	primitiveType GLPrimitiveType
	attribs       []VertexAttrib
	isBound       bool
}

// NewVBO creates a new managed VBO object.
func NewVBO(buffer []Vertex, primitiveType GLPrimitiveType, attribs []VertexAttrib, vboUsage VBOUsage) (*VBO, error) {
	vbo := &VBO{
		bufferUsage:   vboUsage,
		primitiveType: primitiveType,
		attribs:       attribs,
		buffer:        buffer,
	}

	if err := vbo.createVBO(); err != nil {
		return nil, err
	}

	return vbo, nil
}

// Destroy implements a VBO destructor.
func (vbo *VBO) Destroy() {
	gogl.DeleteBuffers(1, &vbo.glVBOID)
}

// Bind the VBO for use and enable registered VertexAttribs.
func (vbo *VBO) Bind() {
	if vbo.isBound {
		core.Log(core.LogErr, "VBO is already bound")
		return
	}

	bufSize := len(vbo.buffer) * VertexByteSize

	gogl.BindBuffer(gogl.ARRAY_BUFFER, vbo.glVBOID)

	if vbo.bufferUsage != StaticDraw {
		gogl.BufferData(gogl.ARRAY_BUFFER, vbo.bufferCapacity, nil, uint32(vbo.bufferUsage))

		if bufSize > 0 {
			gogl.BufferSubData(gogl.ARRAY_BUFFER, 0, bufSize, gogl.Ptr(vbo.buffer))
		}
	}

	for _, attrib := range vbo.attribs {
		gogl.EnableVertexAttribArray(attrib.Index)
		gogl.VertexAttribPointer(attrib.Index, attrib.Size, uint32(attrib.Type), attrib.Normalized, attrib.Stride, gogl.PtrOffset(attrib.Offset))
	}

	vbo.isBound = true
}

// Unbind the VBO.
func (vbo *VBO) Unbind() {
	if vbo.isBound == false {
		core.Log(core.LogErr, "VBO already unbound")
		return
	}

	for _, attrib := range vbo.attribs {
		gogl.DisableVertexAttribArray(attrib.Index)
	}
	gogl.BindBuffer(gogl.ARRAY_BUFFER, 0)
	vbo.isBound = false
}

// Render the VBO using GLDrawArrays.
func (vbo *VBO) Render() {
	if vbo.isBound == false {
		core.Log(core.LogErr, "cannot render unbound VBO")
		return
	}

	gogl.DrawArrays(uint32(vbo.primitiveType), 0, int32(len(vbo.buffer)))
	LogGLErrors()
}

// UpdateBuffer sets the data source for buffer operations.
func (vbo *VBO) UpdateBuffer(buffer []Vertex) error {
	if vbo.isBound {
		return errors.New("cannot update buffer on bound VBO")
	}

	vbo.buffer = buffer

	var bufSize = len(vbo.buffer) * VertexByteSize
	if bufSize > vbo.bufferCapacity || vbo.bufferUsage == StaticDraw {
		if err := vbo.createVBO(); err != nil {
			return err
		}
	}

	return nil
}

// SetAttribs updates the VertexAttribs used by the VBO.
func (vbo *VBO) SetAttribs(attribs []VertexAttrib) error {
	if vbo.isBound {
		return errors.New("cannot update attributes on bound VBO")
	}

	vbo.attribs = attribs

	return nil
}

func (vbo *VBO) createVBO() error {
	var id uint32

	gogl.GenBuffers(1, &id)
	gogl.BindBuffer(gogl.ARRAY_BUFFER, id)
	if gogl.IsBuffer(id) == false {
		return errors.New("failed to generate VBO buffer")
	}

	bufSize := len(vbo.buffer) * VertexByteSize
	bufCap := int(core.NP2(uint(bufSize) + 1))

	if vbo.bufferUsage == StaticDraw && bufSize > 0 {
		gogl.BufferData(gogl.ARRAY_BUFFER, bufSize, gogl.Ptr(vbo.buffer), uint32(vbo.bufferUsage))
	} else {
		gogl.BufferData(gogl.ARRAY_BUFFER, bufCap, nil, uint32(vbo.bufferUsage))
	}

	gogl.BindBuffer(gogl.ARRAY_BUFFER, 0)

	if err := GetGLError(); err != nil {
		gogl.DeleteBuffers(1, &id)
		return err
	}

	vbo.maybeDeleteVBO()

	vbo.glVBOID = id
	vbo.bufferCapacity = bufCap

	return nil
}

func (vbo *VBO) maybeDeleteVBO() {
	if gogl.IsBuffer(vbo.glVBOID) {
		gogl.DeleteBuffers(1, &vbo.glVBOID)
	}

	vbo.glVBOID = 0
}
