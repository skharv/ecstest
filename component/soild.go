package component

import "skharv/ecstest/helper/enum"

type Solid struct {
	Group enum.CollisionGroup // A group of objects to define collision rules.
}

func NewSolid(group enum.CollisionGroup) Solid {
	return Solid{group}
}

func (s Solid) Empty() bool {
	return s.Group == enum.CollisionGroupNone
}
