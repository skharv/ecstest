package system

import (
	"skharv/ecstest/assets"
	"skharv/ecstest/component"
	"skharv/ecstest/helper"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sedyh/mizu/pkg/engine"
)

type Render struct {
	offscreen []*ebiten.Image
}

func NewRender() *Render {
	r := Render{}
	for i := 0; i < helper.RenderLayers; i++ {
		r.offscreen = append(r.offscreen, ebiten.NewImage(helper.ScreenWidth, helper.ScreenHeight))
	}

	return &r
}

func (r *Render) Draw(w engine.World, screen *ebiten.Image) {
	screen.Fill(assets.Background)

	renders := w.View(
		component.Position{},
		component.Render{},
		component.Sprite{},
		component.Hue{},
	).Filter()

	for _, e := range renders {
		var pos *component.Position
		var ren *component.Render
		var spr *component.Sprite
		var hue *component.Hue
		e.Get(&pos, &ren, &spr, &hue)

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(pos.X, pos.Y)
		if hue.Colorful {
			op.ColorM.RotateHue(hue.Value)
		}
		r.offscreen[ren.Value].DrawImage(spr.Image, op)
	}

	op := &ebiten.DrawImageOptions{}
	op.Filter = ebiten.FilterLinear
	for i := 0; i < helper.RenderLayers; i++ {
		screen.DrawImage(r.offscreen[i], op)
		r.offscreen[i].Clear()
	}
}
