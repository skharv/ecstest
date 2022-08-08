package system

import (
	"skharv/ecstest/component"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sedyh/mizu/pkg/engine"
)

type CameraMovement struct{}

func NewCameraMovement() *CameraMovement {
	return &CameraMovement{}
}

func (c *CameraMovement) Update(w engine.World) {
	cameraEntity, found := w.View(
		component.Position{},
		component.Velocity{},
		component.Speed{},
		component.Zoom{},
	).Get()
	if !found {
		return
	}

	var position *component.Position
	var velocity *component.Velocity
	var speed *component.Speed
	var zoom *component.Zoom

	cameraEntity.Get(&position, &velocity, &speed, &zoom)

	X := 0
	Y := 0

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		Y -= int(speed.Value)
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		X -= int(speed.Value)
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		Y += int(speed.Value)
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		X += int(speed.Value)
	}

	velocity.Y = float64(Y)
	velocity.X = float64(X)

	position.X += velocity.X
	position.Y += velocity.Y
}
