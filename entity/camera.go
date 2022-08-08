package entity

import (
	"skharv/ecstest/component"
)

type Camera struct {
	component.Position
	component.Velocity
	component.Speed
	component.Zoom
}
