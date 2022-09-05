package scene

import (
	"skharv/ecstest/assets"
	"skharv/ecstest/component"
	"skharv/ecstest/entity"
	"skharv/ecstest/system"

	"github.com/sedyh/mizu/pkg/engine"
)

type Game struct{}

func (g *Game) Setup(w engine.World) {
	w.AddComponents(
		component.Position{},
		component.Velocity{},
		component.Render{},
		component.Speed{},
		component.Sprite{},
		component.Mouse{},
		component.Hue{},
		component.Size{},
		component.Group{},
		component.Gravity{},
		component.Control{},
		component.Construct{},
	)

	w.AddSystems(
		system.NewRender(),
		system.NewConstruct(),
		system.NewHighlight(),
	)

	w.AddEntities(
		&entity.Construct{
			Construct: component.NewConstruct(10, 10),
		},
		&entity.Cursor{
			Position: component.NewPositionI(0, 0),
			Render:   component.NewRender(20),
			Sprite:   component.NewSprite(assets.Selection),
			Mouse:    component.NewMouse(),
			Hue:      component.NewHue(true, 1),
		},
	)
}
