package component

import "skharv/ecstest/helper"

type Render struct {
	Value int
}

func NewRender(value int) Render {
	if value < 0 {
		value = 0
	} else if value > helper.RenderLayers {
		value = helper.RenderLayers
	}
	return Render{Value: value}
}
