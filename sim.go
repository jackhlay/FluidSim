package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

var conf = Config{
	DynamicColor: false,
	Width:        720,
	Height:       680,

	Particles:  2048,
	Viscosity:  0.9945,
	Turbulence: 0.000000001,

	lowVelocityThreshold: .05,

	Friction:  0.08,
	Drag:      0.045,
	Repulsion: 0.03,
	Bounce:    0.00,
	Gravity:   0.43,

	SpringConstant: .23,
	RestLength:     9,

	Size: 11,
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
