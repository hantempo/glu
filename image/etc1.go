package image

import (
	"fmt"
	"image"
	"image/color"

	glcolor "github.com/hantempo/glu/image/color"
)

const (
	// How many pixels in a block in width or height dimension
	blockWidth = 4
	// How many bytes in a block
	blockSize = 8
)

var codeWordTable = [][]int16{
	{-8, -2, 2, 8},
	{-17, -5, 5, 17},
	{-29, -9, 9, 29},
	{-42, -13, 13, 42},
	{-60, -18, 18, 60},
	{-80, -24, 24, 80},
	{-106, -33, 33, 106},
	{-183, -47, 47, 183},
}

var modifierTableIndex = []uint8{
	2, 3, 1, 0,
}

func extend5to8Bits(v uint8) uint8 {
	v = v & 0x1F
	return (v << 3) | (v >> 2)
}

func extend4to8Bits(v uint8) uint8 {
	v = v & 0x0F
	return (v << 4) | v
}

func clamp(v uint8, modifier int16) uint8 {
	vc := int16(v) + modifier
	if vc < 0 {
		return 0
	} else if vc > 0xFF {
		return 0xFF
	} else {
		return uint8(vc)
	}
}

func calculateSizeETC1(width, height int) int {
	blockSizeWidth := (width + blockWidth - 1) / blockWidth
	blockSizeHeight := (height + blockWidth - 1) / blockWidth
	return blockSizeWidth * blockSizeHeight * blockSize
}

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

	_, yBlockDim := p.BlockDimensions()

	xBlock := x / blockWidth
	yBlock := y / blockWidth

	blockOffset := yBlock*yBlockDim + xBlock
	byteOffset := blockOffset * blockSize

	blockData := p.Pix[byteOffset : byteOffset+8]
	diffBit := blockData[3]&0x02 == 0x01
	flipBit := blockData[3]&0x01 == 0x01

	// pixel offset in this block
	pixelOffsetX, pixelOffsetY := x%blockWidth, y%blockWidth
	pixelOffst := uint(pixelOffsetX*blockWidth + pixelOffsetY)
	var pixelIndex uint8
	if pixelOffst >= 8 {
		pixelIndex = ((blockData[5]>>pixelOffst)&0x01)<<1 + (blockData[7]>>pixelOffst)&0x01
	} else {
		pixelOffst -= 8
		pixelIndex = ((blockData[4]>>pixelOffst)&0x01)<<1 + (blockData[6]>>pixelOffst)&0x01
	}
	fmt.Println(pixelIndex)

	inFirstBlock := true
	if flipBit {
		inFirstBlock = pixelOffsetY >= 2
	} else {
		inFirstBlock = pixelOffsetX >= 2
	}

	var r, g, b, codeWordIndex uint8
	if inFirstBlock {
		if diffBit {
			r = extend5to8Bits((blockData[0] >> 3) & 0x1F)
			g = extend5to8Bits((blockData[1] >> 3) & 0x1F)
			b = extend5to8Bits((blockData[2] >> 3) & 0x1F)
		} else {
			r = extend4to8Bits((blockData[0] >> 4) & 0x0F)
			g = extend4to8Bits((blockData[1] >> 4) & 0x0F)
			b = extend4to8Bits((blockData[2] >> 4) & 0x0F)
		}
		codeWordIndex = (blockData[3] >> 5) & 0x07
	} else {
		if diffBit {
			r = extend5to8Bits((blockData[0]>>3)&0x1F + blockData[0]&0x07)
			b = extend5to8Bits((blockData[1]>>3)&0x1F + blockData[1]&0x07)
			g = extend5to8Bits((blockData[2]>>3)&0x1F + blockData[2]&0x07)
		} else {
			r = extend4to8Bits(blockData[0] & 0x0F)
			g = extend4to8Bits(blockData[1] & 0x0F)
			b = extend4to8Bits(blockData[2] & 0x0F)
		}
		codeWordIndex = (blockData[3] >> 2) & 0x07
	}
	codeWord := codeWordTable[codeWordIndex]
	modifier := codeWord[modifierTableIndex[pixelIndex]]

	return glcolor.RGB{
		clamp(r, modifier),
		clamp(g, modifier),
		clamp(b, modifier),
	}
}

func (p *ETC1) BlockDimensions() (x, y int) {
	x = (p.Rect.Dx() + blockWidth - 1) / blockWidth
	y = (p.Rect.Dy() + blockWidth - 1) / blockWidth
	return
}

func (p *ETC1) Compress(im image.Image) error {
	return nil
}

func (p *ETC1) Uncompress() (image.Image, error) {
	return p, nil
}

func NewETC1(r image.Rectangle) *ETC1 {
	w, h := r.Dx(), r.Dy()
	buf := make([]byte, calculateSizeETC1(w, h))
	return &ETC1{buf, r}
}
