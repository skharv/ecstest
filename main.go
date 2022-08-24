package main

import (
	"log"
	"math/rand"
	"time"

	"skharv/ecstest/scene"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sedyh/mizu/pkg/engine"
)

func main() {
	//assets.Init()
	ebiten.SetWindowSize(800, 600)
	ebiten.SetFPSMode(ebiten.FPSModeVsyncOffMaximum)
	rand.Seed(time.Now().UTC().UnixNano())
	if err := ebiten.RunGame(engine.NewGame(&scene.Game{})); err != nil {
		log.Fatal(err)
	}
}
