package assets

import (
	"embed"
	"image/color"

	"skharv/ecstest/helper"
)

// This is where the images, colors and gradients are loaded.

var (
	Background = color.RGBA{R: 41, G: 44, B: 45, A: 255}
	Acid       = helper.Image(fs, "data/images/enchant-acid-1.png")
	Blob       = helper.Image(fs, "data/images/blob.png")
)

//go:embed "data"
var fs embed.FS
