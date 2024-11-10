package main

import (
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

var conf = Config{
	DynamicColor: false,
	Width:        1280,
	Height:       720,
	Particles:    128,
	Viscosity:    .99,
	Turbulence:   .01,
	Repulsion:    .99,
	Size:         7,
}

func main() {
	// Prepare particles
	particles = make([]Particle, conf.Particles)
	for i := 0; i < conf.Particles; i++ {
		particles[i] = Particle{
			x:  float64(conf.Width / 2),
			y:  float64(conf.Height / 2),
			vx: (2*rand.Float64() - 1),
			vy: -9.81,
		}
	}

	// Initialize the game
	game := &Game{}

	// Set window title and size
	ebiten.SetWindowSize(conf.Width, conf.Height)
	ebiten.SetWindowTitle("Fluid Simulation in Ebiten")

	// Run the game loop
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
