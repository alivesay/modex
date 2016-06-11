package gl

import (
	"errors"
	"github.com/alivesay/modex/core"
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

const initialBufferCapacity = 256

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
	gl.DeleteBuffers(1, &vbo.glVBOID)
}

func (vbo *VBO) Bind() {
	if vbo.isBound {
		core.Log(core.LogErr, "VBO is already bound")
		return
	}

	bufSize := len(vbo.buffer) * VertexByteSize

	gl.BindBuffer(gl.ARRAY_BUFFER, vbo.glVBOID)

	if vbo.bufferUsage != StaticDraw {
		gl.BufferData(gl.ARRAY_BUFFER, vbo.bufferCapacity, nil, uint32(vbo.bufferUsage))

		if bufSize > 0 {
			gl.BufferSubData(gl.ARRAY_BUFFER, 0, bufSize, gl.Ptr(vbo.buffer))
		}
	}

	for _, attrib := range vbo.attribs {
		gl.EnableVertexAttribArray(attrib.Index)
		gl.VertexAttribPointer(attrib.Index, attrib.Size, uint32(attrib.Type), attrib.Normalized, int32(attrib.Stride), gl.PtrOffset(attrib.Offset))

	}

	vbo.isBound = true
}

func (vbo *VBO) Unbind() {
	if vbo.isBound == false {
		core.Log(core.LogErr, "VBO already unbound")
		return
	}

	for _, attrib := range vbo.attribs {
		gl.DisableVertexAttribArray(attrib.Index)
	}
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	vbo.isBound = false
}

func (vbo *VBO) Render() {
	if vbo.isBound == false {
		core.Log(core.LogErr, "cannot render unbound VBO")
		return
	}

	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(vbo.buffer)))
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

	gl.GenBuffers(1, &id)
	gl.BindBuffer(gl.ARRAY_BUFFER, id)
	if gl.IsBuffer(id) == false {
		return errors.New("failed to generate VBO buffer")
	}

	bufSize := len(vbo.buffer) * VertexByteSize
	bufCap := int(core.NP2(uint32(bufSize) + 1))

	switch vbo.bufferUsage {
	case StaticDraw:
		gl.BufferData(gl.ARRAY_BUFFER, bufSize, gl.Ptr(vbo.buffer), uint32(vbo.bufferUsage))
		break
	default:
		gl.BufferData(gl.ARRAY_BUFFER, bufCap, nil, uint32(vbo.bufferUsage))
	}

	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	if err := GetGLError(); err != nil {
		gl.DeleteBuffers(1, &id)
		return err
	}

	vbo.maybeDeleteVBO()

	vbo.glVBOID = id
	vbo.bufferCapacity = bufCap

	return nil
}

func (vbo *VBO) maybeDeleteVBO() {
	if gl.IsBuffer(vbo.glVBOID) {
		gl.DeleteBuffers(1, &vbo.glVBOID)
	}

	vbo.glVBOID = 0
}
