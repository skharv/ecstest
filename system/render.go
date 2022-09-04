package system

import (
	"skharv/ecstest/assets"
	"skharv/ecstest/component"
	"skharv/ecstest/helper"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sedyh/mizu/pkg/engine"
)

type Render struct {
	offscreen *ebiten.Image
}

func NewRender() *Render {
	return &Render{
		offscreen: ebiten.NewImage(helper.ScreenWidth, helper.ScreenHeight),
	}
}

func (r *Render) Draw(w engine.World, screen *ebiten.Image) {
	screen.Fill(assets.Background)

	renders := w.View(
		component.Position{},
		component.Sprite{},
		component.Hue{},
	).Filter()

	for _, e := range renders {
		var pos *component.Position
		var sprite *component.Sprite
		var hue *component.Hue
		e.Get(&pos, &sprite, &hue)

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(pos.X, pos.Y)
		if hue.Colorful {
			op.ColorM.RotateHue(hue.Value)
		}
		r.offscreen.DrawImage(sprite.Image, op)
	}

	op := &ebiten.DrawImageOptions{}
	op.Filter = ebiten.FilterLinear
	screen.DrawImage(r.offscreen, op)
	r.offscreen.Clear()
}
