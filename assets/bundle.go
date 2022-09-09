package assets

import (
	"embed"
	"image/color"

	"skharv/ecstest/helper"
)

// This is where the images, colors and gradients are loaded.

var (
	Background   = color.RGBA{R: 41, G: 44, B: 45, A: 255}
	BlueUnit     = helper.Image(fs, "data/images/blueunit.png")
	RedUnit      = helper.Image(fs, "data/images/redunit.png")
	BlueUnitIcon = helper.Image(fs, "data/images/blueuniticon.png")
	RedUnitIcon  = helper.Image(fs, "data/images/reduniticon.png")
	Tile         = helper.Image(fs, "data/images/tile.png")
	Selection    = helper.Image(fs, "data/images/selection.png")
)

//go:embed "data"
var fs embed.FS
