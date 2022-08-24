package scene

import (
	"skharv/ecstest/assets"
	"skharv/ecstest/component"
	"skharv/ecstest/entity"
	"skharv/ecstest/helper/enum"
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
		component.Solid{},
		component.Gravity{},
		component.Control{},
	)

	w.AddSystems(
		system.NewRender(),
		system.NewPlayer(),
		system.NewVelocity(),
	)

	w.AddEntities(
		&entity.Image{
			Position: component.NewPositionI(0, 0),
			Sprite:   component.NewSprite(assets.Acid),
			Hue:      component.NewHue(false, 0),
		},
		&entity.Image{
			Position: component.NewPositionI(128, 128),
			Sprite:   component.NewSprite(assets.Blob),
			Hue:      component.NewHue(false, 0),
		},
		&entity.Character{
			Position: component.NewPositionI(256, 256),
			Velocity: component.NewVelocityI(0, 0),
			Size:     component.NewSizeI(1, 1),
			Solid:    component.NewSolid(enum.CollisionGroupPlayer),
			Sprite:   component.NewSprite(assets.Blob),
			Gravity:  component.NewGravity(0),
			Control:  component.NewControl(1),
			Hue:      component.NewHue(false, 0),
		},
		&entity.Character{
			Position: component.NewPositionI(100, 100),
			Velocity: component.NewVelocityI(0, 0),
			Size:     component.NewSizeI(1, 1),
			Solid:    component.NewSolid(enum.CollisionGroupPlayer),
			Sprite:   component.NewSprite(assets.Blob),
			Gravity:  component.NewGravity(0),
			Control:  component.NewControl(2),
			Hue:      component.NewHue(true, 600),
		},
	)
}
