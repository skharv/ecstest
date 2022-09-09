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
		component.Button{},
		component.Position{},
		component.Velocity{},
		component.Render{},
		component.Speed{},
		component.Sprite{},
		component.Mouse{},
		component.Hide{},
		component.Hue{},
		component.Ui{},
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
			Hide:     component.NewHide(false),
			Hue:      component.NewHue(true, 0),
		},
		&entity.Icon{
			Position: component.NewPositionI(10, 10),
			Render:   component.NewRender(80),
			Sprite:   component.NewSprite(assets.RedUnitIcon),
			Button:   component.NewButton(enum.UnitRed),
			Size:     component.NewSizeI(32, 32),
			Hide:     component.NewHide(false),
			Hue:      component.NewHue(true, 0),
			Ui:       component.NewUi(),
		},
		&entity.Icon{
			Position: component.NewPositionI(10, 52),
			Render:   component.NewRender(80),
			Sprite:   component.NewSprite(assets.BlueUnitIcon),
			Button:   component.NewButton(enum.UnitBlue),
			Size:     component.NewSizeI(32, 32),
			Hide:     component.NewHide(false),
			Hue:      component.NewHue(true, 0),
			Ui:       component.NewUi(),
		},
	)
}
