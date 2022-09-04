package system

import (
	"skharv/ecstest/assets"
	"skharv/ecstest/component"
	"skharv/ecstest/entity"
	"skharv/ecstest/helper"
	"skharv/ecstest/helper/enum"

	"github.com/sedyh/mizu/pkg/engine"
)

type Construct struct {
}

func NewConstruct() *Construct {
	return &Construct{}
}

func (c *Construct) Update(w engine.World) {
	constructEntity, found := w.View(component.Construct{}).Get()
	if !found {
		return
	}

	var construct *component.Construct

	constructEntity.Get(&construct)

	XOffset := (helper.ScreenWidth - (construct.Width * 64)) / 2
	YOffset := (helper.ScreenHeight - (construct.Height * 64)) / 2

	for i := 0; i < construct.Width; i++ {
		for j := 0; j < construct.Height; j++ {
			w.AddEntities(&entity.Tile{
				Position: component.NewPositionI((64*i)+XOffset, (64*j)+YOffset),
				Sprite:   component.NewSprite(assets.Tile),
				Group:    component.NewGroup(enum.GroupTile),
				Size:     component.NewSizeI(64, 64),
				Hue:      component.NewHue(true, 0),
			})
		}
	}

	w.RemoveEntity(constructEntity)
}
