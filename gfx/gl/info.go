package gl

import (
	"encoding/json"
	"github.com/alivesay/modex/core"
	gogl "github.com/go-gl/gl/all-core/gl"
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
			Vendor:     gogl.GoStr(gogl.GetString(gogl.VENDOR)),
			Version:    gogl.GoStr(gogl.GetString(gogl.VERSION)),
			Renderer:   gogl.GoStr(gogl.GetString(gogl.RENDERER)),
			Extensions: gogl.GoStr(gogl.GetString(gogl.EXTENSIONS)),
		}

		gogl.GetIntegerv(gogl.MAX_TEXTURE_IMAGE_UNITS, &infoInstance.MaxTextureImageUnits)
		gogl.GetIntegerv(gogl.MAX_TEXTURE_SIZE, &infoInstance.MaxTextureSize)
		gogl.GetIntegerv(gogl.MAX_VERTEX_ATTRIBS, &infoInstance.MaxVertexAttribs)
		gogl.GetIntegerv(gogl.MAX_VIEWPORT_DIMS, &infoInstance.MaxViewportDims[0])
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
