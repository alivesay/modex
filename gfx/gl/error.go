package gl

import (
	"errors"
	"strings"

	"github.com/alivesay/modex/core"

	gogl "github.com/go-gl/gl/all-core/gl"
)

var GLErrorStrings = map[uint32]string{
	gogl.NO_ERROR:                      `GL_NO_ERROR`,
	gogl.INVALID_ENUM:                  `GL_INVALID_ENUM`,
	gogl.INVALID_OPERATION:             `GL_INVALID_OPERATION`,
	gogl.INVALID_VALUE:                 `GL_INVALID_VALUE`,
	gogl.INVALID_FRAMEBUFFER_OPERATION: `GL_INVALID_FRAMEBUFFER_OPERATION`,
	gogl.OUT_OF_MEMORY:                 `GL_OUT_OF_MEMORY`,
}

func GetGLError() error {
	errStrings := make([]string, 0)
	for glError := gogl.GetError(); glError != gogl.NO_ERROR; glError = gogl.GetError() {
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
