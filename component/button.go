package component

import "skharv/ecstest/helper/enum"

type Button struct {
	Unit enum.Unit
}

func NewButton(unit enum.Unit) Button {
	return Button{unit}
}

func (b Button) Empty() bool {
	return b.Unit == enum.UnitNone
}
