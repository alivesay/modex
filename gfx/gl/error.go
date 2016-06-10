package gl

import (
	"errors"
	gl "github.com/go-gl/glow/gl"
	"strings"
)

var GLErrorStrings = map[uint32]string{
	gl.NO_ERROR:                      `GL_NO_ERROR`,
	gl.INVALID_ENUM:                  `GL_INVALID_ENUM`,
	gl.INVALID_OPERATION:             `GL_INVALID_OPERATION`,
	gl.INVALID_FRAMEBUFFER_OPERATION: `GL_INVALID_FRAMEBUFFER_OPERATION`,
	gl.OUT_OF_MEMORY:                 `GL_OUT_OF_MEMORY`,
}

func GetGLError() error {
	errStrings := make([]string, 0)
	for glError := gl.GetError(); glError != gl.NO_ERROR; glError = gl.GetError() {
		errStrings = append(errStrings, GLErrorStrings[glError])
	}

	if len(errStrings) > 0 {
		return errors.New(strings.Join(errStrings, ","))
	}

	return nil
}
