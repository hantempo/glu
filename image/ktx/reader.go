package ktx

import (
	"encoding/binary"
	"fmt"
	"image"
	"image/color"
	"io"

	glimage "github.com/hantempo/glu/image"
	glcolor "github.com/hantempo/glu/image/color"
)

type Type int

const (
	GL_BYTE                    Type = 0x00001400
	GL_UNSIGNED_BYTE           Type = 0x00001401
	GL_SHORT                   Type = 0x00001402
	GL_UNSIGNED_SHORT          Type = 0x00001403
	GL_INT                     Type = 0x00001404
	GL_UNSIGNED_INT            Type = 0x00001405
	GL_FLOAT                   Type = 0x00001406
	GL_UNSIGNED_BYTE_3_3_2     Type = 0x00008032
	GL_UNSIGNED_SHORT_4_4_4_4  Type = 0x00008033
	GL_UNSIGNED_SHORT_5_5_5_1  Type = 0x00008034
	GL_UNSIGNED_INT_8_8_8_8    Type = 0x00008035
	GL_UNSIGNED_INT_10_10_10_2 Type = 0x00008036
)

type Format int

const (
	GL_RED             = 0x00001903
	GL_GREEN           = 0x00001904
	GL_BLUE            = 0x00001905
	GL_ALPHA           = 0x00001906
	GL_RGB             = 0x00001907
	GL_RGBA            = 0x00001908
	GL_LUMINANCE       = 0x00001909
	GL_LUMINANCE_ALPHA = 0x0000190A
)

func decodeUint32(buf []byte, isLittleEndianness bool) uint32 {
	if isLittleEndianness {
		return binary.LittleEndian.Uint32(buf[:4])
	} else {
		return binary.BigEndian.Uint32(buf[:4])
	}
}

type decoder struct {
	im            image.Image
	model         color.Model
	width, height int
}

const magic = "\xAB\x4B\x54\x58\x20\x31\x31\xBB\x0D\x0A\x1A\x0A"

func (d *decoder) decode(r io.Reader, configOnly bool) error {
	// Enough to hold the header identifier (12 bytes)
	var tmp [256]byte

	// Read and check the identifier
	buf := tmp[:12]
	if _, err := io.ReadFull(r, buf); err != nil {
		return err
	}
	if magic != string(buf) {
		return fmt.Errorf("KTX reader: invalid identifier [%v]", buf)
	}

	// Read and decide endianness
	buf = tmp[:4]
	if _, err := io.ReadFull(r, buf); err != nil {
		return err
	}
	isLittleEndianness := true
	endianness := decodeUint32(buf, true)
	if endianness == 0x04030201 {
		isLittleEndianness = true
	} else if endianness == 0x01020304 {
		isLittleEndianness = false
	} else {
		return fmt.Errorf("KTX reader: invalid endianness [%v]", buf)
	}

	buf = tmp[:12*4]
	if _, err := io.ReadFull(r, buf); err != nil {
		return err
	}

	glType := decodeUint32(buf[0:4], isLittleEndianness)
	glTypeSize := decodeUint32(buf[4:8], isLittleEndianness)
	glFormat := decodeUint32(buf[8:12], isLittleEndianness)
	glInternalFormat := decodeUint32(buf[12:16], isLittleEndianness)
	glBaseInternalFormat := decodeUint32(buf[16:20], isLittleEndianness)
	fmt.Println(string(glType), glTypeSize, glFormat, glInternalFormat, glBaseInternalFormat)

	d.width = int(decodeUint32(buf[20:24], isLittleEndianness))
	d.height = int(decodeUint32(buf[24:28], isLittleEndianness))

	buf = tmp[:4]
	if _, err := io.ReadFull(r, buf); err != nil {
		return err
	}
	imageSize := int(decodeUint32(buf[0:4], isLittleEndianness))
	fmt.Println(imageSize)

	var (
		gray      *image.Gray
		nrgba4444 *glimage.NRGBA4444
	)
	if glType == uint32(GL_UNSIGNED_BYTE) && glFormat == uint32(GL_LUMINANCE) {
		gray = image.NewGray(image.Rect(0, 0, d.width, d.height))
		if _, err := io.ReadFull(r, gray.Pix); err != nil {
			return err
		}
		d.im = gray
		d.model = color.GrayModel
	} else if glType == uint32(GL_UNSIGNED_SHORT_4_4_4_4) && glFormat == uint32(GL_RGBA) {
		nrgba4444 = glimage.NewNRGBA4444(image.Rect(0, 0, d.width, d.height))
		if _, err := io.ReadFull(r, nrgba4444.Pix); err != nil {
			return err
		}
		d.im = nrgba4444
		d.model = glcolor.NRGBA4444Model
	} else {
		return fmt.Errorf("")
	}

	return nil
}

func Decode(r io.Reader) (image.Image, error) {
	var d decoder
	err := d.decode(r, false)
	return d.im, err
}

func DecodeConfig(r io.Reader) (image.Config, error) {
	var d decoder
	err := d.decode(r, true)
	return image.Config{
		ColorModel: d.model,
		Width:      d.width,
		Height:     d.height,
	}, err
}

func init() {
	image.RegisterFormat("ktx", magic, Decode, DecodeConfig)
}
