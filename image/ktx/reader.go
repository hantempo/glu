package ktx

import (
	"encoding/binary"
	"fmt"
	"image"
	"image/color"
	"io"
	"log"

	"github.com/hantempo/glu/enum"
	glimage "github.com/hantempo/glu/image"
	glcolor "github.com/hantempo/glu/image/color"
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
	log.Printf("glType : %s\n", enum.TypeString(glType))
	log.Printf("glTypeSize : %v\n", glTypeSize)
	log.Printf("glFormat : %s\n", enum.FormatString(glFormat))
	log.Printf("glInternalFormat : %s\n", enum.FormatString(glInternalFormat))
	log.Printf("glBaseInternalFormat : %s\n", enum.FormatString(glBaseInternalFormat))

	d.width = int(decodeUint32(buf[20:24], isLittleEndianness))
	d.height = int(decodeUint32(buf[24:28], isLittleEndianness))
	log.Printf("pixelWidth : %v\n", d.width)
	log.Printf("pixelHeight : %v\n", d.height)
	log.Printf("pixelDepth : %v\n", decodeUint32(buf[28:32], isLittleEndianness))
	log.Printf("numberOfArrayElements : %v\n", decodeUint32(buf[32:], isLittleEndianness))
	log.Printf("numberOfFaces : %v\n", decodeUint32(buf[36:], isLittleEndianness))
	log.Printf("numberOfMipmapLevels : %v\n", decodeUint32(buf[40:], isLittleEndianness))
	log.Printf("numberOfKeyValuePairs : %v\n", decodeUint32(buf[44:], isLittleEndianness))

	buf = tmp[:4]
	if _, err := io.ReadFull(r, buf); err != nil {
		return err
	}
	imageSize := int(decodeUint32(buf[0:4], isLittleEndianness))
	log.Printf("imageSize : %v\n", imageSize)

	var (
		gray      *image.Gray
		nrgba4444 *glimage.NRGBA4444
	)
	if glType == enum.GL_UNSIGNED_BYTE && glFormat == enum.GL_LUMINANCE {
		gray = image.NewGray(image.Rect(0, 0, d.width, d.height))
		if _, err := io.ReadFull(r, gray.Pix); err != nil {
			return err
		}
		d.im = gray
		d.model = color.GrayModel
	} else if glType == enum.GL_UNSIGNED_SHORT_4_4_4_4 && glFormat == enum.GL_RGBA {
		nrgba4444 = glimage.NewNRGBA4444(image.Rect(0, 0, d.width, d.height))
		if _, err := io.ReadFull(r, nrgba4444.Pix); err != nil {
			return err
		}

		d.im = nrgba4444
		d.model = glcolor.NRGBA4444Model
	} else {
		return fmt.Errorf("KTX reader: unsupported type-format combination [%v %v]\n", glType, glFormat)
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
