package system

import (
	"skharv/ecstest/component"
	"skharv/ecstest/helper/controls"
	"skharv/ecstest/helper/num"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sedyh/mizu/pkg/engine"
)

type Player struct{}

func NewPlayer() *Player {
	return &Player{}
}

func (p *Player) Update(w engine.World) {
	players := w.View(
		component.Position{},
		component.Velocity{},
		component.Gravity{},
		component.Sprite{},
		component.Control{},
	).Filter()

	for _, e := range players {
		var pos *component.Position
		var vel *component.Velocity
		var gravity *component.Gravity
		var sprite *component.Sprite
		var control *component.Control
		e.Get(&pos, &vel, &gravity, &sprite, &control)

		e.Get()

		moveDirectionX := 0.0
		if ebiten.IsKeyPressed(controls.Left) {
			moveDirectionX += -1.0
		}
		if ebiten.IsKeyPressed(controls.Right) {
			moveDirectionX += 1.0
		}

		moveDirectionY := 0.0
		if ebiten.IsKeyPressed(controls.Up) {
			moveDirectionY += -1.0
		}
		if ebiten.IsKeyPressed(controls.Down) {
			moveDirectionY += 1.0
		}

		vel.X = num.Lerp(vel.X, control.MoveSpeed*moveDirectionX, control.MoveSpeed)
		vel.Y = num.Lerp(vel.Y, control.MoveSpeed*moveDirectionY, control.MoveSpeed)
	}
}
