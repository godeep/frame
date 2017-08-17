package frame

import (
	"image"
	"image/color"
)

var (
	Black  = image.NewUniform(color.RGBA{0, 0, 0, 255})
	White  = image.NewUniform(color.RGBA{255, 255, 255, 255})
	Yellow = image.NewUniform(color.RGBA{255, 255, 224, 255})
	Green  = image.NewUniform(color.RGBA{0x99, 0xCC, 0x99, 255})
	Red    = image.NewUniform(color.RGBA{0xCC, 0x99, 0x99, 255})
	Gray   = image.NewUniform(color.RGBA{0x55, 0x55, 0x55, 255})
	Mauve  = image.NewUniform(color.RGBA{0x99, 0x99, 0xDD, 255})
)

type Color struct {
	Palette
	Hi Palette
}

var Acme = Color{
	Palette: Palette{Text: Gray, Back: Yellow},
	Hi:      Palette{Text: White, Back: Mauve},
}

type Palette struct {
	Text, Back image.Image
}
