package main

import (
	"image/color"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// Define the game structure
type Game struct{}

// Define Config struct
type Config struct {
	Width      int
	Height     int
	Particles  int
	Viscosity  float64
	Turbulence float64
	Repulsion  float64
	Speed      float64
	Size       float64
}

// Particle struct
type Particle struct {
	x, y   float64 // position
	vx, vy float64 // velocity
}

var particles = make([]Particle, conf.Particles)

// Update runs the game logic each frame
func (g *Game) Update() error {
	// Apply viscosity (friction)
	for i := range particles {
		particles[i].vx *= conf.Viscosity
		particles[i].vy *= conf.Viscosity

		particles[i].vx += (rand.Float64() - 0.5) * conf.Turbulence
		particles[i].vy += (rand.Float64() - 0.5) * conf.Turbulence

		//Move particle
		particles[i].x += particles[i].vx
		particles[i].y += particles[i].vy

		//Edge Detection
		if particles[i].x-conf.Size < float64(conf.Size) || particles[i].x+conf.Size > float64(conf.Width)-conf.Size {
			particles[i].vx = -particles[i].vx
		}
		if particles[i].y-conf.Size < float64(conf.Size) || particles[i].y+conf.Size > float64(conf.Height)-conf.Size {
			particles[i].vy = -particles[i].vy
		}

		for j := range particles {
			if i == j {
				continue
			}
			// Calculate distance between particle[i] and particle[j]
			distance := math.Hypot(particles[i].x-particles[j].x, particles[i].y-particles[j].y)
			if distance < 2*float64(conf.Size) {
				dx := particles[i].x - particles[j].x
				dy := particles[i].y - particles[j].y
				dist := math.Hypot(dx, dy)

				if dist != 0 {
					// Calculate force magnitude based on distance
					forceMagnitude := conf.Repulsion * (2*conf.Size - distance)
					// Normalize the direction of the force and apply it
					particles[i].vx += forceMagnitude * (dx / distance)
					particles[i].vy += forceMagnitude * (dy / distance)
					particles[j].vx -= forceMagnitude * (dx / distance)
					particles[j].vy -= forceMagnitude * (dy / distance)
				}
			}

		}
	}

	return nil
}

// Draw renders each frame
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Clear()
	for _, p := range particles {
		// Use vector to draw filled circles or create your own render function
		vector.DrawFilledCircle(screen, float32(p.x), float32(p.y), float32(conf.Size), color.RGBA{0, 255, 255, 255}, true)
	}

}

// Layout defines the screen size
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return conf.Width, conf.Height // Set the window size (width x height)
}
