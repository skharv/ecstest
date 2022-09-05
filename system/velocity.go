package system

import (
	"skharv/ecstest/component"

	"github.com/sedyh/mizu/pkg/engine"
)

type Velocity struct {
	*component.Velocity
	*component.Position
}

func NewVelocity() *Velocity {
	return &Velocity{}
}

func (v *Velocity) Update(w engine.World) {
	v.Position.Y += v.Velocity.Y
	v.Position.X += v.Velocity.X
}
