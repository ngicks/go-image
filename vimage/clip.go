package vimage

import (
	"image"
	"image/color"
)

var _ image.RGBA64Image = (*Clip)(nil)

type Clip struct {
	Image      image.RGBA64Image
	Limit      image.Rectangle
	Background color.RGBA64
}

func (c *Clip) At(x int, y int) color.Color {
	if !(image.Point{x, y}).In(c.Limit) {
		return c.Background
	}
	return c.Image.At(x, y)
}

func (c *Clip) Bounds() image.Rectangle {
	return c.Limit
}

func (c *Clip) ColorModel() color.Model {
	return c.Image.ColorModel()
}

func (c *Clip) RGBA64At(x int, y int) color.RGBA64 {
	if !(image.Point{x, y}).In(c.Limit) {
		return c.Background
	}
	return c.Image.RGBA64At(x, y)
}
