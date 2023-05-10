package main

import (
	"encoding/binary"
	"hash/fnv"
	"image"
)

func Hash(img image.Image) uint64 {
	hash := fnv.New64a()
	rect := img.Bounds()
	size := (rect.Max.X - rect.Min.X) * (rect.Max.Y - rect.Min.Y) * 16
	data := make([]byte, size)
	pos := 0

	for y := rect.Min.Y; y < rect.Max.Y; y++ {
		for x := rect.Min.X; x < rect.Max.X; x++ {
			color := img.At(x, y)
			r, g, b, a := color.RGBA()

			binary.LittleEndian.PutUint32(data[pos:pos+4], r)
			binary.LittleEndian.PutUint32(data[pos+4:pos+8], g)
			binary.LittleEndian.PutUint32(data[pos+8:pos+12], b)
			binary.LittleEndian.PutUint32(data[pos+12:pos+16], a)

			pos += 16
		}
	}

	hash.Write(data)

	return hash.Sum64()
}
