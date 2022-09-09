package main

import (
	"log"
	"math/rand"
	"time"

	"skharv/ecstest/helper/globals"
	"skharv/ecstest/scene"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sedyh/mizu/pkg/engine"
)

func main() {
	ebiten.SetWindowSize(globals.ScreenWidth, globals.ScreenHeight)
	ebiten.SetFPSMode(ebiten.FPSModeVsyncOffMaximum)
	rand.Seed(time.Now().UTC().UnixNano())
	if err := ebiten.RunGame(engine.NewGame(&scene.Game{})); err != nil {
		log.Fatal(err)
	}
}
