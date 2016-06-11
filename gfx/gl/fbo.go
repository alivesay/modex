package gl

import (
	gles2 "github.com/go-gl/gl/v3.1/gles2"
)

type fbo struct {
	glFBOID uint32
}

func NewFBO() (*fbo, error) {
	fbo := &fbo{}
	gles2.GenFramebuffers(1, &fbo.glFBOID)

	if !gles2.IsFramebuffer(fbo.glFBOID) {
		return nil, GetGLError()
	}

	return fbo, nil
}

func (fbo *fbo) Destroy() {
	gles2.DeleteFramebuffers(1, &fbo.glFBOID)
}

func (fbo *fbo) Bind() {
	gles2.BindFramebuffer(gles2.FRAMEBUFFER, fbo.glFBOID)
}

func (fbo *fbo) Unbind() {
	gles2.BindFramebuffer(gles2.FRAMEBUFFER, 0)
}
