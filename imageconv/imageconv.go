package main

import (
	"flag"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/hantempo/glu/image/ktx"
)

func main() {
	flag.Parse()
	if len(flag.Args()) < 2 {
		return
	}

	input := flag.Arg(0)
	reader, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	im, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	output := flag.Arg(1)
	outputExt := strings.ToUpper(filepath.Ext(output))
	writer, err := os.Create(output)
	if err != nil {
		log.Fatal(err)
	}
	defer writer.Close()

	if outputExt == ".PNG" {
		png.Encode(writer, im)
	} else if outputExt == ".JPEG" || outputExt == ".JPG" {
		jpeg.Encode(writer, im, nil)
	} else {
		log.Fatalf("Unknown output format : %s\n", outputExt)
	}
}
