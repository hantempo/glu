package image

import (
	"encoding/binary"
	"image"
	"image/color"

	glcolor "github.com/hantempo/glu/image/color"
)

// RGB565 is an in-memory image whose At method returns color.RGB565 values.
type RGB565 struct {
	Pix    []uint8
	Stride int
	Rect   image.Rectangle
}

func (p *RGB565) ColorModel() color.Model {
	return glcolor.RGB565Model
}

func (p *RGB565) Bounds() image.Rectangle {
	return p.Rect
}

func (p *RGB565) At(x, y int) color.Color {
	if !(image.Point{x, y}.In(p.Rect)) {
		return glcolor.RGB565{}
	}
	i := p.PixOffset(x, y)
	return glcolor.RGB565{binary.LittleEndian.Uint16(p.Pix[i : i+2])}
}

func (p *RGB565) PixOffset(x, y int) int {
	return (y-p.Rect.Min.Y)*p.Stride + (x-p.Rect.Min.X)*2
}

func (p *RGB565) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	c1 := glcolor.RGB565Model.Convert(c).(glcolor.RGB565)
	binary.LittleEndian.PutUint16(p.Pix[i:i+2], c1.RGB)
}

func NewRGB565(r image.Rectangle) *RGB565 {
	w, h := r.Dx(), r.Dy()
	buf := make([]byte, w*h*2)
	return &RGB565{buf, w * 2, r}
}

// NRGBA4444 is an in-memory image whose At method returns color.NRGBA4444 values.
type NRGBA4444 struct {
	Pix    []byte
	Stride int
	Rect   image.Rectangle
}

func (p *NRGBA4444) ColorModel() color.Model {
	return glcolor.NRGBA4444Model
}

func (p *NRGBA4444) Bounds() image.Rectangle {
	return p.Rect
}

func (p *NRGBA4444) At(x, y int) color.Color {
	if !(image.Point{x, y}.In(p.Rect)) {
		return glcolor.NRGBA4444{}
	}
	i := p.PixOffset(x, y)
	return glcolor.NRGBA4444{binary.LittleEndian.Uint16(p.Pix[i : i+2])}
}

func (p *NRGBA4444) PixOffset(x, y int) int {
	return (y-p.Rect.Min.Y)*p.Stride + (x-p.Rect.Min.X)*2
}

func (p *NRGBA4444) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	c1 := glcolor.NRGBA4444Model.Convert(c).(glcolor.NRGBA4444)
	binary.LittleEndian.PutUint16(p.Pix[i:i+2], c1.Value)
}

func NewNRGBA4444(r image.Rectangle) *NRGBA4444 {
	w, h := r.Dx(), r.Dy()
	buf := make([]byte, w*h*2)
	return &NRGBA4444{buf, w * 2, r}
}
