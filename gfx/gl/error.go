package gl

import (
	"errors"
	"github.com/alivesay/modex/core"
	gles2 "github.com/go-gl/gl/v3.1/gles2"
	"strings"
)

var GLErrorStrings = map[uint32]string{
	gles2.NO_ERROR:                      `GL_NO_ERROR`,
	gles2.INVALID_ENUM:                  `GL_INVALID_ENUM`,
	gles2.INVALID_OPERATION:             `GL_INVALID_OPERATION`,
	gles2.INVALID_VALUE:                 `GL_INVALID_VALUE`,
	gles2.INVALID_FRAMEBUFFER_OPERATION: `GL_INVALID_FRAMEBUFFER_OPERATION`,
	gles2.OUT_OF_MEMORY:                 `GL_OUT_OF_MEMORY`,
}

func GetGLError() error {
	errStrings := make([]string, 0)
	for glError := gles2.GetError(); glError != gles2.NO_ERROR; glError = gles2.GetError() {
		if errString, ok := GLErrorStrings[glError]; ok {
			errStrings = append(errStrings, errString)
		} else {
			errStrings = append(errStrings, "UnhandledGLError")
		}
	}

	if len(errStrings) > 0 {
		return errors.New(strings.Join(errStrings, ","))
	}

	return nil
}

func LogGLErrors() {
	if err := GetGLError(); err != nil {
		core.Log(core.LogErr, err)
	}
}
