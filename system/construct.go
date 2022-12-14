package system

import (
	"skharv/ecstest/assets"
	"skharv/ecstest/component"
	"skharv/ecstest/entity"
	"skharv/ecstest/helper/enum"
	"skharv/ecstest/helper/globals"

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

	XOffset := (globals.ScreenWidth - (construct.Width * 64)) / 2
	YOffset := (globals.ScreenHeight - (construct.Height * 64)) / 2

	for i := 0; i < construct.Width; i++ {
		for j := 0; j < construct.Height; j++ {
			w.AddEntities(&entity.Tile{
				Position: component.NewPositionI((64*i)+XOffset, (64*j)+YOffset),
				Render:   component.NewRender(10),
				Sprite:   component.NewSprite(assets.Tile),
				Group:    component.NewGroup(enum.GroupTile),
				Size:     component.NewSizeI(64, 64),
				Hide:     component.NewHide(false),
				Hue:      component.NewHue(true, 0),
			})
		}
	}

	w.RemoveEntity(constructEntity)
}
