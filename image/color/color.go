package color

import "image/color"

// NGrayAlpha represents a 16-bit non-alpha-premultiplied color,
// having 8 bits for each of grayscale and alpha.
type NGrayAlpha struct {
	G, A uint8
}

func (c NGrayAlpha) RGBA() (r, g, b, a uint32) {
	as := uint32(c.A)
	r = uint32(c.G) * 0x101 * as / 0xFF
	g = r
	b = r
	a = as * 0x101
	return
}

// R8 represents a 8-bit opaque color with only red channel
type R8 struct {
	R uint8
}

func (c R8) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R) * 0x101
	g = 0x0000
	b = 0x0000
	a = 0xFFFF
	return
}

// RGB565 represents a 16-bit opaque color,
// having 5 bits for red, blue and 6 bits for green.
type RGB565 struct {
	RGB uint16
}

func (c RGB565) RGBA() (r, g, b, a uint32) {
	r = ((uint32(c.RGB) >> 11) & 0x1F) * 0xFFFF / 0x1F
	g = ((uint32(c.RGB) >> 5) & 0x3F) * 0xFFFF / 0x3F
	b = ((uint32(c.RGB) >> 0) & 0x1F) * 0xFFFF / 0x1F
	a = 0xFFFF
	return
}

// NRGBA4444 represents a 16-bit non-alpha-premultiplied color,
// having 4 bits for each of red, green, blue and alpha.
type NRGBA4444 struct {
	Value uint16
}

func (c NRGBA4444) RGBA() (r, g, b, a uint32) {
	a4 := uint32(c.Value & 0xF)
	r = ((uint32(c.Value) >> 12) & 0x0F) * 0x1111 * a4 / 0xF
	g = ((uint32(c.Value) >> 8) & 0x0F) * 0x1111 * a4 / 0xF
	b = ((uint32(c.Value) >> 4) & 0x0F) * 0x1111 * a4 / 0xF
	a = a4 * 0x1111
	return
}

// NRGBA5551 represents a 16-bit non-alpha-premultiplied color,
// having 5 bits for each of red, green, blue and 1 bit for alpha.
type NRGBA5551 struct {
	Value uint16
}

func (c NRGBA5551) RGBA() (r, g, b, a uint32) {
	a1 := uint32(c.Value & 0x01)
	if a1 != 0 {
		r = ((uint32(c.Value) >> 11) & 0x1F) * 0xFFFF / 0x1F
		g = ((uint32(c.Value) >> 6) & 0x1F) * 0xFFFF / 0x1F
		b = ((uint32(c.Value) >> 1) & 0x1F) * 0xFFFF / 0x1F
		a = ((uint32(c.Value) >> 0) & 0x01) * 0xFFFF
	}
	return
}

// NBGRA8888 represents a 64-bit non-alpha-premultiplied color,
// having 8 bits for each of red, green, blue and alpha.
type NBGRA8888 struct {
	b, g, r, a uint8
}

func (c NBGRA8888) RGBA() (r, g, b, a uint32) {
	a8 := uint32(c.a)
	r = uint32(c.r) * 0x101 * a8 / 0xFF
	g = uint32(c.g) * 0x101 * a8 / 0xFF
	b = uint32(c.b) * 0x101 * a8 / 0xFF
	a = a8 * 0x101
	return
}

// Models for GL color types
var (
	NGrayAlphaModel color.Model = color.ModelFunc(nGrayAlphaModel)
	R8Model         color.Model = color.ModelFunc(r8Model)
	RGB565Model     color.Model = color.ModelFunc(rgb565Model)
	NRGBA4444Model  color.Model = color.ModelFunc(nRGBA4444Model)
	NRGBA5551Model  color.Model = color.ModelFunc(nRGBA5551Model)
	NBGRA8888Model  color.Model = color.ModelFunc(nBGRA8888Model)
)

func nGrayAlphaModel(c color.Color) color.Color {
	if _, ok := c.(NGrayAlpha); ok {
		return c
	}

	r, _, _, a := c.RGBA()
	return NGrayAlpha{uint8((r * 0xFFFF / a) >> 8), uint8(a >> 8)}
}

func r8Model(c color.Color) color.Color {
	if _, ok := c.(R8); ok {
		return c
	}

	r, _, _, _ := c.RGBA()
	return R8{uint8(r & 0xFF)}
}

func rgb565Model(c color.Color) color.Color {
	if _, ok := c.(RGB565); ok {
		return c
	}

	r, g, b, _ := c.RGBA()
	rs := uint16(r * 0x1F / 0xFFFF)
	gs := uint16(g * 0x3F / 0xFFFF)
	bs := uint16(b * 0x1F / 0xFFFF)
	return RGB565{uint16(rs<<11 | gs<<5 | bs)}
}

func nRGBA4444Model(c color.Color) color.Color {
	if _, ok := c.(NRGBA4444); ok {
		return c
	}

	r, g, b, a := c.RGBA()
	if a == 0xFFFF {
		return NRGBA4444{uint16(r&0xF000 | g&0x0F00 | b&0x00F0 | 0x000F)}
	} else if a == 0x0000 {
		return NRGBA4444{uint16(0x0000)}
	} else {
		return NRGBA4444{uint16((r*0xFFFF/a)&0xF000 | (g*0xFFFF/a)&0x0F00 | (b*0xFFFF/a)&0x00F0 | a&0x000F)}
	}
}

func nRGBA5551Model(c color.Color) color.Color {
	if _, ok := c.(NRGBA5551); ok {
		return c
	}

	r, g, b, a := c.RGBA()
	rs := uint16(r * 0x1F / 0xFFFF)
	gs := uint16(g * 0x1F / 0xFFFF)
	bs := uint16(b * 0x1F / 0xFFFF)
	as := uint16(a & 1)
	return NRGBA5551{uint16(rs<<11 | gs<<6 | bs<<1 | as)}
}

func nBGRA8888Model(c color.Color) color.Color {
	if _, ok := c.(NBGRA8888); ok {
		return c
	}

	r, g, b, a := c.RGBA()
	return NBGRA8888{uint8((b * 0xFFFF / a) >> 8), uint8((g * 0xFFFF / a) >> 8), uint8((r * 0xFFFF / a) >> 8), uint8(a >> 8)}
}
