package entity

import (
	"skharv/ecstest/component"
)

type Tile struct {
	component.Position
	component.Render
	component.Sprite
	component.Group
	component.Size
	component.Hide
	component.Hue
}
