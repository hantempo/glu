package color

import (
	"image/color"
	"testing"
)

func TestNGrayAlpha(t *testing.T) {
	var c NGrayAlpha
	if r, g, b, a := c.RGBA(); r != 0 || g != 0 || b != 0 || a != 0 {
		t.Errorf("r=0x%X g=0x%X b=0x%X a=0x%X", r, g, b, a)
	}

	c = NGrayAlpha{0x55, 0xAA}
	if r, g, b, a := c.RGBA(); r != 0x38E3 || g != 0x38E3 || b != 0x38E3 || a != 0xAAAA {
		t.Errorf("r=0x%X g=0x%X b=0x%X a=0x%X", r, g, b, a)
	}

	c = NGrayAlphaModel.Convert(color.RGBA64{0x38E3, 0x38E3, 0x38E3, 0xAAAA}).(NGrayAlpha)
	if cnew := (NGrayAlpha{0x55, 0xAA}); c != cnew {
		t.Error()
	}
}

func TestR8(t *testing.T) {
	var c R8
	if r, g, b, a := c.RGBA(); r != 0 || g != 0 || b != 0 || a != 0xFFFF {
		t.Errorf("r=0x%X g=0x%X b=0x%X a=0x%X", r, g, b, a)
	}

	c = R8{uint8(0x55)}
	if r, g, b, a := c.RGBA(); r != 0x5555 || g != 0x0000 || b != 0x0000 || a != 0xFFFF {
		t.Errorf("r=0x%X g=0x%X b=0x%X a=0x%X", r, g, b, a)
	}

	c = R8Model.Convert(color.RGBA64{0x5151, 0x5151, 0x5151, 0xFFFF}).(R8)
	if cnew := (R8{0x51}); c != cnew {
		t.Error()
	}
}

func TestRGB(t *testing.T) {
	var c RGB
	if r, g, b, a := c.RGBA(); r != 0 || g != 0 || b != 0 || a != 0xFFFF {
		t.Errorf("r=0x%X g=0x%X b=0x%X a=0x%X", r, g, b, a)
	}

	c = RGB{0x5F, 0x2E, 0x3C}
	if r, g, b, a := c.RGBA(); r != 0x5F5F || g != 0x2E2E || b != 0x3C3C || a != 0xFFFF {
		t.Errorf("r=0x%X g=0x%X b=0x%X a=0x%X", r, g, b, a)
	}

	c = RGBModel.Convert(color.RGBA64{0x5151, 0xB6B6, 0x5151, 0xFFFF}).(RGB)
	if cnew := (RGB{0x51, 0xB6, 0x51}); c != cnew {
		t.Error()
	}
}

func TestRGB565(t *testing.T) {
	var c RGB565
	if r, g, b, a := c.RGBA(); r != 0 || g != 0 || b != 0 || a != 0xFFFF {
		t.Errorf("r=0x%X g=0x%X b=0x%X a=0x%X", r, g, b, a)
	}

	c = RGB565{uint16(0x55AA)}
	if r, g, b, a := c.RGBA(); r != 0x5294 || g != 0xB6DA || b != 0x5294 || a != 0xFFFF {
		t.Errorf("r=0x%X g=0x%X b=0x%X a=0x%X", r, g, b, a)
	}

	c = RGB565Model.Convert(color.RGBA64{0x5151, 0xB6B6, 0x5151, 0xFFFF}).(RGB565)
	if cnew := (RGB565{0x4D89}); c != cnew {
		t.Error()
	}
}

func BenchmarkRGB565(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c := RGB565{uint16(i)}
		r, g, b, a := c.RGBA()
		RGB565Model.Convert(color.RGBA64{uint16(r), uint16(g), uint16(b), uint16(a)})
	}
}

func TestNRGBA4444(t *testing.T) {
	var c NRGBA4444
	if r, g, b, a := c.RGBA(); r != 0 || g != 0 || b != 0 || a != 0 {
		t.Errorf("r=0x%X g=0x%X b=0x%X a=0x%X", r, g, b, a)
	}

	c = NRGBA4444{uint16(0x55AA)}
	if r, g, b, a := c.RGBA(); r != 0x38E3 || g != 0x38E3 || b != 0x71C6 || a != 0xAAAA {
		t.Errorf("r=0x%X g=0x%X b=0x%X a=0x%X", r, g, b, a)
	}

	c = NRGBA4444Model.Convert(color.RGBA64{0x38E3, 0x38E3, 0x71C6, 0xAAAA}).(NRGBA4444)
	if cnew := (NRGBA4444{0x55AA}); c != cnew {
		t.Error()
	}
}

func TestNRGBA5551(t *testing.T) {
	var c NRGBA5551
	if r, g, b, a := c.RGBA(); r != 0 || g != 0 || b != 0 || a != 0 {
		t.Errorf("r=0x%X g=0x%X b=0x%X a=0x%X", r, g, b, a)
	}

	c = NRGBA5551{uint16(0x55AB)}
	if r, g, b, a := c.RGBA(); r != 0x5294 || g != 0xB5AC || b != 0xAD6A || a != 0xFFFF {
		t.Errorf("r=0x%X g=0x%X b=0x%X a=0x%X", r, g, b, a)
	}

	c = NRGBA5551{uint16(0x55AA)}
	if r, g, b, a := c.RGBA(); r != 0 || g != 0 || b != 0 || a != 0 {
		t.Errorf("r=0x%X g=0x%X b=0x%X a=0x%X", r, g, b, a)
	}

	c = NRGBA5551Model.Convert(color.RGBA64{0x5555, 0xAAAA, 0x5555, 0xFFFF}).(NRGBA5551)
	if cnew := (NRGBA5551{0x5515}); c != cnew {
		t.Errorf("rgba=0x%X", c.Value)
	}
}

func TestNBGRA8888(t *testing.T) {
	var c NBGRA8888
	if r, g, b, a := c.RGBA(); r != 0 || g != 0 || b != 0 || a != 0 {
		t.Errorf("r=0x%X g=0x%X b=0x%X a=0x%X", r, g, b, a)
	}

	c = NBGRA8888{0x11, 0x22, 0x33, 0x44}
	if r, g, b, a := c.RGBA(); r != 0x0DA7 || g != 0x091A || b != 0x048D || a != 0x4444 {
		t.Errorf("r=0x%X g=0x%X b=0x%X a=0x%X", r, g, b, a)
	}

	c = NBGRA8888Model.Convert(color.RGBA64{0x0DA7, 0x091A, 0x048D, 0x4444}).(NBGRA8888)
	if cnew := (NBGRA8888{0x11, 0x22, 0x33, 0x44}); c != cnew {
		t.Error()
	}
}
