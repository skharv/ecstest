package system

import (
	"skharv/ecstest/assets"
	"skharv/ecstest/component"
	"skharv/ecstest/helper/globals"
	"sort"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sedyh/mizu/pkg/engine"
)

type Render struct {
	offscreen []*ebiten.Image
}

func NewRender() *Render {
	r := Render{}
	for i := 0; i < globals.RenderLayers; i++ {
		r.offscreen = append(r.offscreen, ebiten.NewImage(globals.ScreenWidth, globals.ScreenHeight))
	}

	return &r
}

func (r *Render) Draw(w engine.World, screen *ebiten.Image) {
	screen.Fill(assets.Background)

	renders := w.View(
		component.Position{},
		component.Render{},
		component.Sprite{},
		component.Hide{},
		component.Hue{},
	).Filter()

	var populatedLayers []int

	for _, e := range renders {
		var pos *component.Position
		var ren *component.Render
		var spr *component.Sprite
		var hid *component.Hide
		var hue *component.Hue
		e.Get(&pos, &ren, &spr, &hid, &hue)

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(pos.X, pos.Y)
		if hue.Colorful {
			op.ColorM.RotateHue(hue.Value)
		}
		if !hid.Value {
			layerFound := false

			for _, layer := range populatedLayers {
				if layer == ren.Value {
					layerFound = true
				}
			}
			if !layerFound {
				populatedLayers = append(populatedLayers, ren.Value)
			}

			r.offscreen[ren.Value].DrawImage(spr.Image, op)
		}
	}

	sort.Slice(populatedLayers, func(i, j int) bool {
		return populatedLayers[i] < populatedLayers[j]
	})

	op := &ebiten.DrawImageOptions{}
	op.Filter = ebiten.FilterLinear
	for _, v := range populatedLayers {
		screen.DrawImage(r.offscreen[v], op)
		r.offscreen[v].Clear()
	}
}
