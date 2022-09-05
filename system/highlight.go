package system

import (
	"skharv/ecstest/component"
	"skharv/ecstest/helper/enum"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sedyh/mizu/pkg/engine"
)

type Highlight struct {
	*component.Mouse
	*component.Position
}

func NewHighlight() *Highlight {
	return &Highlight{}
}

func (h *Highlight) Update(w engine.World) {
	tiles := w.View(
		component.Position{},
		component.Group{},
		component.Size{},
	).Filter()

	for _, t := range tiles {
		var pos *component.Position
		var grp *component.Group
		var siz *component.Size
		t.Get(&pos, &grp, &siz)

		if grp.Group != enum.GroupTile {
			continue
		}

		x, y := ebiten.CursorPosition()

		if float64(x) > pos.X && float64(x) < pos.X+siz.W {
			if float64(y) > pos.Y && float64(y) < pos.Y+siz.H {
				h.Position.X = pos.X
				h.Position.Y = pos.Y
			}
		}
	}
}
