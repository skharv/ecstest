package entity

import "skharv/ecstest/component"

type Character struct {
	component.Position
	component.Velocity
	component.Size
	component.Solid
	component.Sprite
	component.Gravity
	component.Control
	component.Hue
}
