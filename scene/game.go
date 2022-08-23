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
		component.Zoom{},
		component.Sprite{},
		component.Hue{},
	)

	w.AddSystems(
		&system.CameraMovement{},
		&system.Render{},
	)
	w.AddEntities(
		&entity.Camera{},
		&entity.Image{},
	)
}
