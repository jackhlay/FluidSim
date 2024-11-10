package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

var conf = Config{
	DynamicColor: false,
	Width:        720,
	Height:       480,
	Particles:    32768,
	Viscosity:    0.98,
	Turbulence:   0.1,
	Repulsion:    0.13,
	Bounce:       0.09,
	Gravity:      .73,
	Size:         3.0,
}

func main() {

	// Set window title and size
	ebiten.SetWindowSize(conf.Width, conf.Height)
	ebiten.SetWindowTitle("Fluid Sim")
	initParticles()

	// Initialize the game
	game := &Game{}

	// Run the game loop
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
