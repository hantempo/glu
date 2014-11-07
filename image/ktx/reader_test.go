package ktx

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"strings"
	"testing"

	glcolor "github.com/hantempo/glu/image/color"
)

var goodTestData = []struct {
	input  []byte
	dims   image.Rectangle
	model  color.Model
	output []color.Color
}{
	{
		input: []byte{
			'\xAB', 'K', 'T', 'X', ' ', '1', '1', '\xBB', '\r', '\n', '\x1A', '\n',
			1, 2, 3, 4, // litter endian
			0x01, 0x14, 0x00, 0x00, // glType=GL_UNSIGNED_BYTE
			0x01, 0x00, 0x00, 0x00, // glTypeSize=1
			0x09, 0x19, 0x00, 0x00, // glFormat=GL_LUMINANCE
			0x40, 0x80, 0x00, 0x00, // glInternalFormat=GL_LUMINANCE8
			0x09, 0x19, 0x00, 0x00, // glBaseInternalFormat=GL_LUMINANCE
			0x04, 0x00, 0x00, 0x00, // width=4,
			0x01, 0x00, 0x00, 0x00, // height=1,
			0x00, 0x00, 0x00, 0x00, // depth=0,
			0x00, 0x00, 0x00, 0x00, // numberOfArrayElements=0
			0x00, 0x00, 0x00, 0x00, // numberOfFaces=0
			0x01, 0x00, 0x00, 0x00, // numberOfMipmapLevels=1
			0x00, 0x00, 0x00, 0x00, // numberOfKeyValuePairs=0,
			0x04, 0x00, 0x00, 0x00, // imageSize=4,
			0x5A, 0xA5, 0x2B, 0xB2, // imageData
		},
		dims:  image.Rect(0, 0, 4, 1),
		model: color.GrayModel,
		output: []color.Color{
			color.Gray{0x5A},
			color.Gray{0xA5},
			color.Gray{0x2B},
			color.Gray{0xB2},
		},
	},
	{
		input: []byte{
			'\xAB', 'K', 'T', 'X', ' ', '1', '1', '\xBB', '\r', '\n', '\x1A', '\n',
			1, 2, 3, 4, // litter endian
			0x33, 0x80, 0x00, 0x00, // glType=GL_UNSIGNED_SHORT_4_4_4_4
			0x02, 0x00, 0x00, 0x00, // glTypeSize=2
			0x08, 0x19, 0x00, 0x00, // glFormat=GL_RGBA
			0x56, 0x80, 0x00, 0x00, // glInternalFormat=GL_RGBA4
			0x08, 0x19, 0x00, 0x00, // glBaseInternalFormat=GL_RGBA
			0x02, 0x00, 0x00, 0x00, // width=2,
			0x01, 0x00, 0x00, 0x00, // height=1,
			0x00, 0x00, 0x00, 0x00, // depth=0,
			0x00, 0x00, 0x00, 0x00, // numberOfArrayElements=0
			0x00, 0x00, 0x00, 0x00, // numberOfFaces=0
			0x01, 0x00, 0x00, 0x00, // numberOfMipmapLevels=1
			0x00, 0x00, 0x00, 0x00, // numberOfKeyValuePairs=0,
			0x04, 0x00, 0x00, 0x00, // imageSize=4,
			0x5A, 0xA5, 0x2B, 0xB2, // imageData
		},
		dims:  image.Rect(0, 0, 2, 1),
		model: glcolor.NRGBA4444Model,
		output: []color.Color{
			glcolor.NRGBA4444{0xA55A},
			glcolor.NRGBA4444{0xB22B},
		},
	},
}

func TestDecode(t *testing.T) {
	for _, test := range goodTestData {
		im, err := Decode(bytes.NewReader(test.input))
		if err != nil {
			t.Fatal(err)
		}
		if im != nil {
			dims := im.Bounds()
			if dims != test.dims {
				t.Errorf("Wrong image size : expected(%v) got(%v)\n", test.dims, dims)
			}

			pixels := test.output
			for i := dims.Min.X; i < dims.Max.X; i++ {
				for j := dims.Min.Y; j < dims.Max.Y; j++ {
					offset := j*dims.Dx() + i
					pixel := im.At(i, j)
					expectedPixel := pixels[offset]
					fmt.Printf("%t %t\n", pixel, expectedPixel)
					if im.At(i, j) != expectedPixel {
						t.Errorf("Wrong pixel at [%v %v] : %v\n", i, j, pixel)
					}
				}
			}
		} else {
			t.Error("No image produced")
		}
	}
}

func TestDecodeConfig(t *testing.T) {
	for _, test := range goodTestData {
		config, err := DecodeConfig(bytes.NewReader(test.input))
		if err != nil {
			t.Fatal(err)
		}

		if config.Width != test.dims.Dx() || config.Height != test.dims.Dy() {
			t.Errorf("Wrong image size : expected(%vx%v) got(%vx%v)\n", test.dims.Dx(), test.dims.Dy(), config.Width, config.Height)
		}

		if config.ColorModel != test.model {
			t.Errorf("Wrong color model")
		}
	}
}

var badTestData = []struct {
	input []byte
	err   string
}{
	{
		input: []byte{
			'\xAB', 'K', 'T', 'X', '1', '1', '1', '\xBB', '\r', '\n', '\x1A', '\n',
		},
		err: "KTX reader: invalid identifier",
	},
	{
		input: []byte{
			'\xAB', 'K', 'T', 'X', ' ', '1', '1', '\xBB', '\r', '\n', '\x1A', '\n',
			1, 1, 1, 1,
		},
		err: "KTX reader: invalid endianness",
	},
}

func TestDecodeError(t *testing.T) {
	for _, test := range badTestData {
		_, err := Decode(bytes.NewReader(test.input))

		if err == nil {
			t.Errorf("Expected pattern of error message (%s), got no error", test.err)
		} else if strings.Contains(err.Error(), test.err) == false {
			t.Errorf("Expected pattern of error message (%s), got (%s)", test.err, err.Error())
		}
	}

}
