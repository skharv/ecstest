package system

import (
	"skharv/ecstest/component"

	"github.com/sedyh/mizu/pkg/engine"
)

type Gravity struct {
	*component.Velocity
	*component.Gravity
}

func NewGravity() *Gravity {
	return &Gravity{}
}

func (g *Gravity) Update(_ engine.World) {
	g.Velocity.Y += g.Gravity.Value
}
