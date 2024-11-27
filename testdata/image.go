package testdata

import (
	"embed"
	"image"
	"image/png"
	"sync"
)

//go:embed image
var imageDataFs embed.FS

var imageMap sync.Map // map[string]image.Image

func RequestImage(name string) image.Image {
	i, ok := imageMap.Load(name)
	if ok {
		return i.(image.Image)
	}
	f, err := imageDataFs.Open("image/" + name + ".drawio.png")
	if err != nil {
		panic(err)
	}
	defer func() { _ = f.Close() }()
	img, err := png.Decode(f)
	if err != nil {
		panic(err)
	}
	i, _ = imageMap.LoadOrStore(name, img)
	return i.(image.Image)
}
