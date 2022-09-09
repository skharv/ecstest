package component

import "skharv/ecstest/helper/globals"

type Render struct {
	Value int
}

func NewRender(value int) Render {
	if value < 0 {
		value = 0
	} else if value > globals.RenderLayers {
		value = globals.RenderLayers
	}
	return Render{Value: value}
}
