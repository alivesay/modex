package pixbuf

type Bitmapable interface {
	Width() uint
	SetWidth(uint)
	Height() uint
	SetHeight(uint)
	Data() []byte
	SetData([]byte)
	Get(uint, uint)
}

type Bitmap struct {
	Bitmapable
	data []byte
}

func NewBitmap() *Bitmap {
	return &Bitmap{}
}
