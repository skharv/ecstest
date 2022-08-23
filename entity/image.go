package entity

import (
	"skharv/ecstest/component"
)

type Image struct {
	component.Position
	component.Sprite
	component.Hue
}
