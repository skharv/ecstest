package entity

import "skharv/ecstest/component"

type Character struct {
	component.Position
	component.Velocity
	component.Size
	component.Group
	component.Sprite
	component.Gravity
	component.Control
	component.Hue
}
