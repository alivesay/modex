package gl

import (
	gl "github.com/go-gl/glow/gl"
)

type fbo struct {
	glFBOID uint32
}

func NewFBO() (*fbo, error) {
	fbo := &fbo{}
	gl.GenFramebuffers(1, &fbo.glFBOID)

	if !gl.IsFramebuffer(fbo.glFBOID) {
		return nil, GetGLError()
	}

	return fbo, nil
}

func (fbo *fbo) Destroy() {
	gl.DeleteFramebuffers(1, &fbo.glFBOID)
}

func (fbo *fbo) Bind() {
	gl.BindFramebuffer(gl.FRAMEBUFFER, fbo.glFBOID)
}

func (fbo *fbo) Unbind() {
	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
}
