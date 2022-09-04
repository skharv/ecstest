package scene

import (
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
		component.Speed{},
		component.Sprite{},
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
	)

	w.AddEntities(
		&entity.Construct{
			Construct: component.NewConstruct(10, 10),
		},
	)
}
