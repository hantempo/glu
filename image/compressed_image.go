package image

import "image"

type BlockCompressedImage interface {
	image.Image
	Compress(im image.Image) error
	Uncompress() (image.Image, error)
	BlockDimensions() (int, int)
}
