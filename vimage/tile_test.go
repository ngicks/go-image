package vimage

import (
	"image"
	"image/png"
	"os"
	"testing"

	"github.com/ngicks/go-image/testdata"
)

func TestTile(t *testing.T) {
	img := testdata.RequestImage("letter")

	tile := &Tile{
		Image: ConvertImage(img),
	}

	clipped := &Clip{
		Image: tile,
		Limit: image.Rect(0, 0, 500, 500),
	}

	i := HintOpaque(clipped, false)

	f, err := os.Create("./tile.out.png")
	if err != nil {
		panic(err)
	}
	defer func() { _ = f.Close() }()
	err = png.Encode(f, i)
	if err != nil {
		panic(err)
	}
}
