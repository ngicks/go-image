package vimage

import (
	"image"
	"image/color"
	"math"
)

var _ image.RGBA64Image = (*Tile)(nil)

type Tile struct {
	Image image.RGBA64Image
}

func (t *Tile) translate(x, y int) (int, int) {
	bounds := t.Image.Bounds()

	// think as if things start at 0,0.
	// then add start point offset at last.

	x = x % bounds.Dx()
	y = y % bounds.Dy()

	if x < 0 {
		x += bounds.Max.X
	}
	if y < 0 {
		y += bounds.Max.Y
	}

	return x + bounds.Min.X, y + bounds.Min.Y
}

func (t *Tile) At(x int, y int) color.Color {
	return t.Image.At(t.translate(x, y))
}

func (t *Tile) Bounds() image.Rectangle {
	return image.Rectangle{
		Min: image.Point{X: math.MinInt, Y: math.MinInt},
		Max: image.Point{X: math.MaxInt, Y: math.MaxInt},
	}
}

func (t *Tile) ColorModel() color.Model {
	return t.Image.ColorModel()
}

func (t *Tile) RGBA64At(x int, y int) color.RGBA64 {
	return t.Image.RGBA64At(t.translate(x, y))
}
