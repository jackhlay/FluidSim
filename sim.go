package main

import (
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

var conf = Config{
	Width:      1280,
	Height:     720,
	Particles:  1024,
	Viscosity:  .99,
	Turbulence: .03,
	Repulsion:  .5,
	Speed:      -9.81,
	Size:       7,
}

func main() {
	//Prepare particles
	for i := 0; i < conf.Particles; i++ {
		particles = append(particles, Particle{
			x:  float64(conf.Width / 2),
			y:  float64(conf.Height / 2),
			vx: (2*rand.Float64() - 1),
			vy: 0,
		})
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
