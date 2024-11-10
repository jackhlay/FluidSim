package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

var conf = Config{
	DynamicColor:   true,
	Width:          720,
	Height:         640,
	Particles:      8192,
	Viscosity:      0.99,
	Turbulence:     0.00001,
	Repulsion:      0.29,
	Bounce:         0.75,
	Gravity:        0.3,
	SpringConstant: 0.025,
	RestLength:     5,
	Size:           3,
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
