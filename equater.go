//go:build !ent

package main

import (
	"image"
)

type SimpleImageEquater struct{}

func (s *SimpleImageEquater) Equal(left, right image.Image) bool {
	leftBounds := left.Bounds()
	rightBounds := right.Bounds()

	if !leftBounds.Eq(rightBounds) {
		return false
	}

	minPoint := leftBounds.Min
	maxPoint := leftBounds.Max

	for y := minPoint.Y; y < maxPoint.Y; y++ {
		for x := minPoint.X; x < maxPoint.X; x++ {
			if !ColorEqual(left.At(x, y), right.At(x, y)) {
				return false
			}
		}
	}

	return true
}

func NewSimpleImageEquater() *SimpleImageEquater {
	return &SimpleImageEquater{}
}
