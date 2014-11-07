package main

import (
	"flag"
	"fmt"
	"image"
	"os"

	_ "github.com/hantempo/glu/image/ktx"
)

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		return
	}

	input := flag.Arg(0)
	reader, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	defer reader.Close()

	im, imfmt, err := image.Decode(reader)
	fmt.Printf("%v %T %v\n", im.Bounds(), im.ColorModel(), imfmt)
}
