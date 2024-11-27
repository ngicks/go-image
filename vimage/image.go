package vimage

import (
	"image"
	"image/color"
	"iter"
)

var _ image.RGBA64Image = (*rgba64Image[image.Image])(nil)

func ConvertImage[T image.Image](i T) image.RGBA64Image {
	if i64, ok := any(i).(image.RGBA64Image); ok {
		return i64
	}
	return &rgba64Image[T]{Image: i}
}

type rgba64Image[T image.Image] struct {
	Image T
}

func (r *rgba64Image[T]) At(x int, y int) color.Color {
	return r.Image.At(x, y)
}

func (r *rgba64Image[T]) Bounds() image.Rectangle {
	return r.Image.Bounds()
}

func (r *rgba64Image[T]) ColorModel() color.Model {
	return r.Image.ColorModel()
}

func (r *rgba64Image[T]) RGBA64At(x int, y int) color.RGBA64 {
	return color.RGBA64Model.Convert(r.Image.At(x, y)).(color.RGBA64)
}

func EnumeratePix(rect image.Rectangle, rowWise bool) iter.Seq2[int, int] {
	if rowWise {
		return func(yield func(int, int) bool) {
			for y := rect.Min.Y; y < rect.Max.Y; y++ {
				for x := rect.Min.X; x < rect.Max.X; x++ {
					if !yield(x, y) {
						return
					}
				}
			}
		}
	} else {
		// col-wise
		return func(yield func(int, int) bool) {
			for x := rect.Min.X; x < rect.Max.X; x++ {
				for y := rect.Min.Y; y < rect.Max.Y; y++ {
					if !yield(x, y) {
						return
					}
				}
			}
		}
	}
}
