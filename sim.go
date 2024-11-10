package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

var conf = Config{
	DynamicColor:   false,
	Width:          1600,
	Height:         900,
	Particles:      4096,
	Viscosity:      0.9045,
	Turbulence:     0.000000001,
	Repulsion:      0.03,
	Bounce:         0.035,
	Gravity:        0.43,
	SpringConstant: .23,
	RestLength:     6,
	Size:           7,
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
