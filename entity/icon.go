package entity

import (
	"skharv/ecstest/component"
)

type Icon struct {
	component.Position
	component.Render
	component.Sprite
	component.Button
	component.Size
	component.Hide
	component.Hue
	component.Ui
}
