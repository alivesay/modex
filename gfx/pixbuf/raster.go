package pixbuf

type RasterCursor struct {
	X int
	Y int
}

func (cursor *RasterCursor) Reset() {
	cursor.X = 0
	cursor.Y = 0
}

type Raster struct {
	fgColor RGBA32
	bgColor RGBA32
}
