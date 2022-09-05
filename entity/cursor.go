package entity

import "skharv/ecstest/component"

type Cursor struct {
	component.Position
	component.Render
	component.Sprite
	component.Mouse
	component.Hide
	component.Hue
}
