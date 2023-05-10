//go:build !ent

package main

import (
	"image"
)

type SimpleImageEquater struct{}

func (s *SimpleImageEquater) Equal(left, right image.Image) bool {
	return Hash(left) == Hash(right)
}

func NewSimpleImageEquater() *SimpleImageEquater {
	return &SimpleImageEquater{}
}
