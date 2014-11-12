package image

import (
	"image"
	"image/color"
	"testing"
)

type testImage interface {
	image.Image
	Set(int, int, color.Color)
}

func cmp(t *testing.T, cm color.Model, c0, c1 color.Color) bool {
	r0, g0, b0, a0 := cm.Convert(c0).RGBA()
	r1, g1, b1, a1 := cm.Convert(c1).RGBA()
	return r0 == r1 && g0 == g1 && b0 == b1 && a0 == a1
}

func TestImage(t *testing.T) {
	testImage := []testImage{
		NewRGB565(image.Rect(0, 0, 10, 10)),
		NewNRGBA4444(image.Rect(0, 0, 10, 10)),
	}
	for _, m := range testImage {
		if !image.Rect(0, 0, 10, 10).Eq(m.Bounds()) {
			t.Errorf("%T: want bounds %v, got %v", m, image.Rect(0, 0, 10, 10), m.Bounds())
			continue
		}
		if !cmp(t, m.ColorModel(), color.Transparent, m.At(6, 3)) {
			t.Errorf("%T: at (6, 3), want a zero color, got %v", m, m.At(6, 3))
			continue
		}
		m.Set(6, 3, color.Opaque)
		if !cmp(t, m.ColorModel(), color.Opaque, m.At(6, 3)) {
			t.Errorf("%T: at (6, 3), want a non-zero color, got %v", m, m.At(6, 3))
			continue
		}
	}
}

func TestCompressedImage(t *testing.T) {
	testImage := []CompressedImage{
		NewETC1(image.Rect(0, 0, 10, 10)),
	}
	for _, m := range testImage {
		if !image.Rect(0, 0, 10, 10).Eq(m.Bounds()) {
			t.Errorf("%T: want bounds %v, got %v", m, image.Rect(0, 0, 10, 10), m.Bounds())
			continue
		}
	}
}
