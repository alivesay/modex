package gl

import (
	"errors"

	"github.com/alivesay/modex/core"
	gles2 "github.com/go-gl/gl/v3.1/gles2"
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
	StaticDraw  VBOUsage = gles2.STATIC_DRAW
	DynamicDraw VBOUsage = gles2.DYNAMIC_DRAW
	StreamDraw  VBOUsage = gles2.STREAM_DRAW
)

type VBO struct {
	glVBOID        uint32
	bufferUsage    VBOUsage
	bufferCapacity int
	bufferAttribs  []VertexAttrib
	// TODO: support vbo doublebuffering
	//currentBuffer uint32
	buffer  []Vertex
	attribs []VertexAttrib
	isBound bool
}

func NewVBO(buffer []Vertex, attribs []VertexAttrib, vboUsage VBOUsage) (*VBO, error) {
	vbo := &VBO{
		bufferUsage: vboUsage,
		attribs:     attribs,
		buffer:      buffer,
	}

	if err := vbo.createVBO(); err != nil {
		return nil, err
	}

	return vbo, nil
}

func (vbo *VBO) Destroy() {
	gles2.DeleteBuffers(1, &vbo.glVBOID)
}

func (vbo *VBO) Bind() {
	if vbo.isBound {
		core.Log(core.LogErr, "VBO is already bound")
		return
	}

	bufSize := len(vbo.buffer) * VertexByteSize

	gles2.BindBuffer(gles2.ARRAY_BUFFER, vbo.glVBOID)

	if vbo.bufferUsage != StaticDraw {
		gles2.BufferData(gles2.ARRAY_BUFFER, vbo.bufferCapacity, nil, uint32(vbo.bufferUsage))

		if bufSize > 0 {
			gles2.BufferSubData(gles2.ARRAY_BUFFER, 0, bufSize, gles2.Ptr(vbo.buffer))
		}
	}

	for _, attrib := range vbo.attribs {
		gles2.EnableVertexAttribArray(attrib.Index)
		gles2.VertexAttribPointer(attrib.Index, attrib.Size, uint32(attrib.Type), attrib.Normalized, int32(attrib.Stride), gles2.PtrOffset(attrib.Offset))

	}

	vbo.isBound = true
}

func (vbo *VBO) Unbind() {
	if vbo.isBound == false {
		core.Log(core.LogErr, "VBO already unbound")
		return
	}

	for _, attrib := range vbo.attribs {
		gles2.DisableVertexAttribArray(attrib.Index)
	}
	gles2.BindBuffer(gles2.ARRAY_BUFFER, 0)
	vbo.isBound = false
}

func (vbo *VBO) Render() {
	if vbo.isBound == false {
		core.Log(core.LogErr, "cannot render unbound VBO")
		return
	}

	gles2.DrawArrays(gles2.TRIANGLES, 0, int32(len(vbo.buffer)))
	LogGLErrors()
}

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

func (vbo *VBO) UpdateAttribs(attribs []VertexAttrib) error {
	if vbo.isBound {
		return errors.New("cannot update attributes on bound VBO")
	}

	vbo.attribs = attribs

	return nil
}

func (vbo *VBO) createVBO() error {
	var id uint32

	gles2.GenBuffers(1, &id)
	gles2.BindBuffer(gles2.ARRAY_BUFFER, id)
	if gles2.IsBuffer(id) == false {
		return errors.New("failed to generate VBO buffer")
	}

	bufSize := len(vbo.buffer) * VertexByteSize
	bufCap := int(core.NP2(uint32(bufSize) + 1))

	switch vbo.bufferUsage {
	case StaticDraw:
		if bufSize > 0 {
			gles2.BufferData(gles2.ARRAY_BUFFER, bufSize, gles2.Ptr(vbo.buffer), uint32(vbo.bufferUsage))
		}
		break
	default:
		gles2.BufferData(gles2.ARRAY_BUFFER, bufCap, nil, uint32(vbo.bufferUsage))
	}

	gles2.BindBuffer(gles2.ARRAY_BUFFER, 0)

	if err := GetGLError(); err != nil {
		gles2.DeleteBuffers(1, &id)
		return err
	}

	vbo.maybeDeleteVBO()

	vbo.glVBOID = id
	vbo.bufferCapacity = bufCap

	return nil
}

func (vbo *VBO) maybeDeleteVBO() {
	if gles2.IsBuffer(vbo.glVBOID) {
		gles2.DeleteBuffers(1, &vbo.glVBOID)
	}

	vbo.glVBOID = 0
}
