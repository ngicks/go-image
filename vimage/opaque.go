package vimage

import (
	"image"
	"image/color"
)

var _ image.RGBA64Image = (*Opaque[*image.RGBA64])(nil)

// Opaque is a virtual image wrapper that hints opaqueness of the image.
type Opaque[T image.RGBA64Image] struct {
	image  T
	opaque bool
}

// CacheOpaque scans entire image to cache whether it is an opaque image or not.
func CacheOpaque[T image.RGBA64Image](i T) *Opaque[T] {
	if op, ok := any(i).(interface{ Opaque() bool }); ok {
		return &Opaque[T]{
			image:  i,
			opaque: op.Opaque(),
		}
	}
	for x, y := range EnumeratePix(i.Bounds(), true) {
		if i.RGBA64At(x, y).A != 0xffff {
			return &Opaque[T]{
				image:  i,
				opaque: false,
			}
		}
	}
	return &Opaque[T]{
		image:  i,
		opaque: true,
	}
}

// HintOpaque hints opaqueness of i.
// Callers must have prior knowledge of opaqueness and
// put it correctly
// A correct hint avoid full-scan of image when encoding.
func HintOpaque[T image.RGBA64Image](i T, opaque bool) *Opaque[T] {
	return &Opaque[T]{
		image:  i,
		opaque: opaque,
	}
}

func (o *Opaque[T]) At(x int, y int) color.Color {
	return o.image.At(x, y)
}

func (o *Opaque[T]) Bounds() image.Rectangle {
	return o.image.Bounds()
}

func (o *Opaque[T]) ColorModel() color.Model {
	return o.image.ColorModel()
}

func (o *Opaque[T]) RGBA64At(x int, y int) color.RGBA64 {
	return o.image.RGBA64At(x, y)
}

func (o *Opaque[T]) Opaque() bool {
	return o.opaque
}
