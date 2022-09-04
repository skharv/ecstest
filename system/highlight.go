package system

import "skharv/ecstest/component"

type Highlight struct {
	*component.Position
	*component.Group
	*component.Size
}

func NewHighlight() *Highlight {
	return &Highlight{}
}
