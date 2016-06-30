package gl

import (
	gogl "github.com/go-gl/gl/all-core/gl"
)

// TODO: make public export
type fbo struct {
	glFBOID uint32
}

func NewFBO() (*fbo, error) {
	fbo := &fbo{}

	gogl.GenFramebuffers(1, &fbo.glFBOID)

	if !gogl.IsFramebuffer(fbo.glFBOID) {
		return nil, GetGLError()
	}

	return fbo, nil
}

func (fbo *fbo) Destroy() {
	gogl.DeleteFramebuffers(1, &fbo.glFBOID)
}

func (fbo *fbo) Bind() {
	gogl.BindFramebuffer(gogl.FRAMEBUFFER, fbo.glFBOID)
}

func (fbo *fbo) Unbind() {
	gogl.BindFramebuffer(gogl.FRAMEBUFFER, 0)
}
