package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

var imageEquater ImageEquater

// ImageEquater is an interface that provides a method to assess the equality
// between two images.
type ImageEquater interface {
	// Equal returns true if both images are equal, false otherwise.
	Equal(image.Image, image.Image) bool
}

// ColorEqual returns true if the R, G, B, and A components of both provided
// Color values match.
func ColorEqual(left, right color.Color) bool {
	lR, lG, lB, lA := left.RGBA()
	rR, rG, rB, rA := right.RGBA()

	return lR == rR && lG == rG && lB == rB && lA == rA
}

func loadImage(filePath string) (image.Image, error) {
	r, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	return png.Decode(r)
}

func main() {
	args := os.Args

	if len(args) < 3 {
		fmt.Fprintln(os.Stderr, "missing arguments - specify two filepaths to assess equality")
		os.Exit(2)
	}

	leftImage, err := loadImage(args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to parse image", args[1])
		os.Exit(2)
	}

	rightImage, err := loadImage(args[2])
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to parse image", args[2])
		os.Exit(2)
	}

	if imageEquater == nil {
		imageEquater = NewSimpleImageEquater()
	}

	if !imageEquater.Equal(leftImage, rightImage) {
		fmt.Println("images differ")
		os.Exit(1)
	}
}
