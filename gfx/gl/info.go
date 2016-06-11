package gl

import (
	"encoding/json"
	"github.com/alivesay/modex/core"
	gles2 "github.com/go-gl/gl/v3.1/gles2"
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
			Vendor:     gles2.GoStr(gles2.GetString(gles2.VENDOR)),
			Version:    gles2.GoStr(gles2.GetString(gles2.VERSION)),
			Renderer:   gles2.GoStr(gles2.GetString(gles2.RENDERER)),
			Extensions: gles2.GoStr(gles2.GetString(gles2.EXTENSIONS)),
		}

		gles2.GetIntegerv(gles2.MAX_TEXTURE_IMAGE_UNITS, &infoInstance.MaxTextureImageUnits)
		gles2.GetIntegerv(gles2.MAX_TEXTURE_SIZE, &infoInstance.MaxTextureSize)
		gles2.GetIntegerv(gles2.MAX_VERTEX_ATTRIBS, &infoInstance.MaxVertexAttribs)
		gles2.GetIntegerv(gles2.MAX_VIEWPORT_DIMS, &infoInstance.MaxViewportDims[0])
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
