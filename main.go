package main

import (
	"flag"
	"fmt"
	"image"
	"image/png"
	"os"
	"path"

	"github.com/nfnt/resize"
	"github.com/pkg/errors"
)

func main() {
	xcassets := ""
	source := ""
	flag.StringVar(&source, "source", "", "Source image location")
	flag.StringVar(&xcassets, "xcassets", "", "location of the xcassets directory")
	flag.Parse()
	if xcassets == "" || source == "" {
		flag.Usage()
		os.Exit(0)
	}
	manifest, err := loadManifest(xcassets)
	if err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}

	for _, img := range manifest.Images {
		size := img.GetSizeValue()
		scale := img.GetScaleValue()
		sourceImage, err := loadSourceImage(source)
		newSize := uint(size * scale)
		fmt.Println("Create", img.Filename, "as", newSize, "x", newSize)
		newImage := resize.Resize(newSize, newSize, sourceImage, resize.Lanczos3)
		out, err := os.Create(path.Join(xcassets, img.Filename))
		if err != nil {
			panic(err)
		}
		err = png.Encode(out, newImage)
		if err != nil {
			panic(err)
		}
		out.Close()

	}

}
func loadSourceImage(source string) (image.Image, error) {
	existingImageFile, err := os.Open(source)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to load source image")
	}
	defer existingImageFile.Close()

	// Alternatively, since we know it is a png already
	// we can call png.Decode() directly
	loadedImage, err := png.Decode(existingImageFile)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to decode source image")
	}
	return loadedImage, nil
}

func loadManifest(xcassets string) (*Manifest, error) {
	manifestpath := path.Join(xcassets, "Contents.json")

	fmt.Println("Reading manifest from:", manifestpath)

	file, err := os.Open(manifestpath)
	if err != nil {
		return nil, errors.WithMessage(err, "xcassets Path Invalid")
	}
	manifest := &Manifest{}
	err = manifest.Parse(file)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to parse manifest")
	}
	return manifest, nil
}
