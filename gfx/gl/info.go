package gl

import (
	"encoding/json"
	"github.com/alivesay/modex/core"
	gl "github.com/go-gl/glow/gl"
	"sync"
)

type Info struct {
	MaxTextureImageUnits int32
	MaxTextureSize       int32
	MaxVertexAttribs     int32
	MaxViewportDims      [2]int32
	Vendor               string
	Version              string
	Renderer             string
	Extensions           string
}

var infoInstance *Info
var infoOnce sync.Once

func GetInstanceInfo() *Info {
	infoOnce.Do(func() {
		infoInstance = &Info{
			Vendor:     gl.GoStr(gl.GetString(gl.VENDOR)),
			Version:    gl.GoStr(gl.GetString(gl.VERSION)),
			Renderer:   gl.GoStr(gl.GetString(gl.RENDERER)),
			Extensions: gl.GoStr(gl.GetString(gl.EXTENSIONS)),
		}

		gl.GetIntegerv(gl.MAX_TEXTURE_IMAGE_UNITS, &infoInstance.MaxTextureImageUnits)
		gl.GetIntegerv(gl.MAX_TEXTURE_SIZE, &infoInstance.MaxTextureSize)
		gl.GetIntegerv(gl.MAX_VERTEX_ATTRIBS, &infoInstance.MaxVertexAttribs)
		gl.GetIntegerv(gl.MAX_VIEWPORT_DIMS, &infoInstance.MaxViewportDims[0])
	})
	return infoInstance
}

func (info *Info) String() string {
	out, err := json.Marshal(info)
	if err != nil {
		core.Log(core.LogErr, err)
		return ""
	}

	return string(out)
}
