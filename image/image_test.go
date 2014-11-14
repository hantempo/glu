package image

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"testing"

	glcolor "github.com/hantempo/glu/image/color"
)

type testImage interface {
	image.Image
	Set(int, int, color.Color)
}

func cmp(cm color.Model, c0, c1 color.Color) bool {
	r0, g0, b0, a0 := cm.Convert(c0).RGBA()
	r1, g1, b1, a1 := cm.Convert(c1).RGBA()
	return r0 == r1 && g0 == g1 && b0 == b1 && a0 == a1
}

func TestImage(t *testing.T) {
	testImage := []testImage{
		NewRGB565(image.Rect(0, 0, 10, 10)),
		NewRGB(image.Rect(0, 0, 10, 10)),
		NewNRGBA4444(image.Rect(0, 0, 10, 10)),
	}
	for _, m := range testImage {
		if !image.Rect(0, 0, 10, 10).Eq(m.Bounds()) {
			t.Errorf("%T: want bounds %v, got %v", m, image.Rect(0, 0, 10, 10), m.Bounds())
			continue
		}
		if !cmp(m.ColorModel(), color.Transparent, m.At(6, 3)) {
			t.Errorf("%T: at (6, 3), want a zero color, got %v", m, m.At(6, 3))
			continue
		}
		m.Set(6, 3, color.Opaque)
		if !cmp(m.ColorModel(), color.Opaque, m.At(6, 3)) {
			t.Errorf("%T: at (6, 3), want a non-zero color, got %v", m, m.At(6, 3))
			continue
		}
	}
}

func TestCompressedImage(t *testing.T) {
	dims := image.Rect(0, 0, 10, 10)
	testImage := []BlockCompressedImage{
		NewETC1(dims),
	}
	testBlockDims := [][2]int{
		{3, 3},
	}
	for i, m := range testImage {
		if dims != m.Bounds() {
			t.Errorf("%T: want bounds %v, got %v", m, image.Rect(0, 0, 10, 10), m.Bounds())
			continue
		}

		bx, by := testBlockDims[i][0], testBlockDims[i][1]
		bxx, byy := m.BlockDimensions()
		if bx != bxx || by != byy {
			t.Errorf("%T: want block dimensions %vx%x, got %vx%x", m, bx, by, bxx, byy)
			continue
		}

		//// Compress and uncompress a black image
		//blackImage := image.NewRGBA(dims)
		//for j := 0; j < 10; j++ {
		//for i := 0; i < 10; i++ {
		//blackImage.Set(i, j, color.Black)
		//}
		//}
		//{
		//err := m.Compress(blackImage)
		//if err != nil {
		//t.Error(err)
		//}
		//uncomImage, err := m.Uncompress()
		//if err != nil {
		//t.Error(err)
		//}
		//for j := 0; j < 10; j++ {
		//for i := 0; i < 10; i++ {
		//if !cmp(color.RGBAModel, color.Black, uncomImage.At(i, j)) {
		//t.Errorf("%T: at (%v, %v), want a black color, got %v", m, i, j, uncomImage.At(i, j))
		//continue
		//}
		//}
		//}
		//}

		//// Compress and uncompress a white image
		//whiteImage := image.NewRGBA(dims)
		//for j := 0; j < 10; j++ {
		//for i := 0; i < 10; i++ {
		//whiteImage.Set(i, j, color.White)
		//}
		//}
		//{
		//err := m.Compress(whiteImage)
		//if err != nil {
		//t.Error(err)
		//}
		//uncomImage, err := m.Uncompress()
		//if err != nil {
		//t.Error(err)
		//}
		//for j := 0; j < 10; j++ {
		//for i := 0; i < 10; i++ {
		//if !cmp(color.RGBAModel, color.White, uncomImage.At(i, j)) {
		//t.Errorf("%T: at (%v, %v), want a white color, got %v", m, i, j, uncomImage.At(i, j))
		//continue
		//}
		//}
		//}
		//}
	}

	im := NewRGB(image.Rect(0, 0, 2, 1))
	im.Set(0, 0, glcolor.RGB{0xF0, 0xA2, 0x15})
	im.Set(1, 0, glcolor.RGB{0xF4, 0xA6, 0x19})
	writer, _ := os.Create("/work/sandbox/etcpack/test.png")
	defer writer.Close()
	png.Encode(writer, im)
}
