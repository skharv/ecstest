package system

import (
	"skharv/ecstest/component"
	"skharv/ecstest/helper/controls"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sedyh/mizu/pkg/engine"
)

type Player struct {
	*component.Position
	*component.Velocity
	*component.Gravity
	*component.Sprite
	*component.Control
}

func NewPlayer() *Player {
	return &Player{}
}

func (p *Player) Update(w engine.World) {
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

	p.Velocity.X = p.Control.MoveSpeed * moveDirectionX
	p.Velocity.Y = p.Control.MoveSpeed * moveDirectionY
}
