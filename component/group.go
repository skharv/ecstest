package component

import "skharv/ecstest/helper/enum"

type Group struct {
	Group enum.Group
}

func NewGroup(group enum.Group) Group {
	return Group{group}
}

func (g Group) Empty() bool {
	return g.Group == enum.GroupNone
}
