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
	ebiten.SetWindowSize(640, 640)
	ebiten.SetWindowSizeLimits(64, 64, -1, -1)
	ebiten.SetFPSMode(ebiten.FPSModeVsyncOffMaximum)
	ebiten.SetWindowResizable(false)
	rand.Seed(time.Now().UTC().UnixNano())
	if err := ebiten.RunGame(engine.NewGame(&scene.Game{})); err != nil {
		log.Fatal(err)
	}
}
