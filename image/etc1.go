package image

import (
	"image"
	"image/color"

	glcolor "github.com/hantempo/glu/image/color"
)

//const (
//blockWidth = 4
//)

//func calculateSizeETC1(width, height int) int {
//blockSizeWidth = ((width + blockWidth - 1)
//}

type ETC1 struct {
	Pix  []uint8
	Rect image.Rectangle
}

func (p *ETC1) ColorModel() color.Model {
	return glcolor.RGBModel
}

func (p *ETC1) Bounds() image.Rectangle {
	return p.Rect
}

func (p *ETC1) At(x, y int) color.Color {
	if !(image.Point{x, y}.In(p.Rect)) {
		return glcolor.RGB{}
	}
	return glcolor.RGB{}
}

func NewETC1(r image.Rectangle) *ETC1 {
	w, h := r.Dx(), r.Dy()
	buf := make([]byte, w*h*2)
	return &ETC1{buf, r}
}
