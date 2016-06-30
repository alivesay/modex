package gl

import (
	"errors"
	"image"
	"image/draw"

	"github.com/alivesay/modex/core"
	"github.com/alivesay/modex/gfx/prim"

	gogl "github.com/go-gl/gl/all-core/gl"
)

type TextureTarget uint32

const (
	Texture2D      TextureTarget = gogl.TEXTURE_2D
	TextureCubeMap               = gogl.TEXTURE_CUBE_MAP
)

type TextureParamName uint32

const (
	TextureMinFilter TextureParamName = gogl.TEXTURE_MIN_FILTER
	TextureMagFilter                  = gogl.TEXTURE_MAG_FILTER
	TextureWrapS                      = gogl.TEXTURE_WRAP_S
	TextureWrapT                      = gogl.TEXTURE_WRAP_T
	// TextureWrapR				   = gogl.TEXTURE_WRAP_R
)

type TextureParamValue uint32

const (
	ClampToEdge    TextureParamValue = gogl.CLAMP_TO_EDGE
	Repeat                           = gogl.REPEAT
	MirroredRepeat                   = gogl.MIRRORED_REPEAT
	Nearest                          = gogl.NEAREST
	Linear                           = gogl.LINEAR

//	NearestMipMapNearest                   = gogl.NEAREST_MIPMAP_NEAREST
//	NearestMipMapLinear                    = gogl.NEAREST_MIPMAP_LINEAR
//	LinearMipMapNearest                    = gogl.LINEAR_MIPMAP_NEAREST
//	LinearMipMapLinear                     = gogl.LINEAR_MIPMAP_LINEAR
)

type TextureParameter struct {
	Name  TextureParamName
	Value TextureParamValue
}

type Texture struct {
	prim.Dimensions
	glTextureID uint32
	target      TextureTarget
	params      []TextureParameter
}

func NewTextureFromImage(src image.Image, target TextureTarget, params []TextureParameter) (*Texture, error) {
	w := src.Bounds().Dx()
	h := src.Bounds().Dy()

	rgba := image.NewNRGBA(src.Bounds())
	if rgba.Stride != rgba.Rect.Size().X*4 {
		return nil, errors.New("unsupported image stride")
	}

	draw.Draw(rgba, rgba.Bounds(), src, image.Point{0, 0}, draw.Src)

	texSize := int32(core.NP2(uint(core.MAX(w, h))))

	tex, err := NewTexture(w, h, target, params)
	if err != nil {
		return nil, err
	}

	tex.Bind()

	for _, param := range tex.params {
		if err := tex.checkParameter(param); err != nil {
			return nil, err
		}

		gogl.TexParameteri(uint32(tex.target), uint32(param.Name), int32(param.Value))
	}

	gogl.TexImage2D(uint32(tex.target),
		0,                  // mipmap level
		gogl.RGBA8,         // internal format
		texSize,            // width
		texSize,            // height
		0,                  // border
		gogl.RGBA,          // format
		gogl.UNSIGNED_BYTE, // type
		gogl.Ptr(rgba.Pix)) // data

	tex.Unbind()

	LogGLErrors()

	return tex, nil
}

func NewTexture(width, height int, target TextureTarget, params []TextureParameter) (*Texture, error) {
	maxTextureSize := int(GetInstanceInfo().MaxTextureSize)

	if width > maxTextureSize || height > maxTextureSize {
		return nil, errors.New("GL_MAX_TEXTURE_SIZE exceeded")
	}

	tex := &Texture{
		target: target,
		params: params,
	}

	gogl.GenTextures(1, &tex.glTextureID)
	gogl.ActiveTexture(gogl.TEXTURE0)
	gogl.BindTexture(uint32(tex.target), tex.glTextureID)
	if !gogl.IsTexture(tex.glTextureID) {
		return nil, GetGLError()
	}
	gogl.BindTexture(uint32(tex.target), 0)

	LogGLErrors()

	return tex, nil
}

func (tex *Texture) Destroy() {
	gogl.DeleteTextures(1, &tex.glTextureID)
}

func (tex *Texture) Bind() {
	gogl.ActiveTexture(gogl.TEXTURE0)
	gogl.BindTexture(gogl.TEXTURE_2D, tex.glTextureID)
}

func (tex *Texture) Unbind() {
	gogl.BindTexture(uint32(tex.target), 0)
}

func (tex *Texture) checkParamMinFilter(param TextureParameter) error {
	switch param.Value {
	case Nearest, Linear: //, NearestMipMapNearest, NearestMipMapLinear, LinearMipMapNearest, LinearMipMapLinear:
		tex.params = append(tex.params, param)
		return nil
	}
	return errors.New("invalid MinFilter value")
}

func (tex *Texture) checkParamMagFilter(param TextureParameter) error {
	switch param.Value {
	case Nearest, Linear:
		tex.params = append(tex.params, param)
		return nil
	}

	return errors.New("invalid MagFilter value")
}

func (tex *Texture) checkParamWrapSFilter(param TextureParameter) error {
	switch param.Value {
	case ClampToEdge, Repeat, MirroredRepeat:
		tex.params = append(tex.params, param)
		return nil
	}
	return errors.New("invalid MagFilter value")
}

func (tex *Texture) checkParamWrapTFilter(param TextureParameter) error {
	switch param.Value {
	case ClampToEdge, Repeat, MirroredRepeat:
		tex.params = append(tex.params, param)
		return nil
	}
	return errors.New("invalid MagFilter value")
}

func (tex *Texture) checkParameter(param TextureParameter) error {
	switch param.Name {
	case TextureMinFilter:
		return tex.checkParamMinFilter(param)
	case TextureMagFilter:
		return tex.checkParamMagFilter(param)
	case TextureWrapS:
		return tex.checkParamWrapSFilter(param)
	case TextureWrapT:
		return tex.checkParamWrapTFilter(param)
	}

	return errors.New("invalid texture parameter name")
}

func (tex *Texture) GLTextureID() uint32 {
	return tex.glTextureID
}
