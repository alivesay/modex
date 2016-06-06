package gl

import (
	"encoding/json"

	"github.com/alivesay/modex/core"
	gl "github.com/go-gl/glow/gl"
)

type Info struct {
	MaxTextureImageUnits int32
	MaxTextureSize       int32
	MaxVertexAttribs     int32
	Vendor               string
	Version              string
	Renderer             string
	Extensions           string
}

func NewInfo() *Info {
	info := &Info{
		Vendor:     gl.GoStr(gl.GetString(gl.VENDOR)),
		Version:    gl.GoStr(gl.GetString(gl.VERSION)),
		Renderer:   gl.GoStr(gl.GetString(gl.RENDERER)),
		Extensions: gl.GoStr(gl.GetString(gl.EXTENSIONS)),
	}

	gl.GetIntegerv(gl.MAX_TEXTURE_IMAGE_UNITS, &info.MaxTextureImageUnits)
	gl.GetIntegerv(gl.MAX_TEXTURE_SIZE, &info.MaxTextureSize)
	gl.GetIntegerv(gl.MAX_VERTEX_ATTRIBS, &info.MaxVertexAttribs)

	return info
}

func (info *Info) String() string {
	out, err := json.Marshal(info)
	if err != nil {
		core.Log(core.LOG_ERR, err)
		return ""
	}

	return string(out)
}