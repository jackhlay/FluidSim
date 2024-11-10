package main

import (
	"fmt"
	"image/color"

	"time"

	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var startTime = time.Now()

// Game structure
type Game struct{}

// Config struct for simulation parameters
type Config struct {
	DynamicColor bool
	Width        int
	Height       int
	Particles    int
	Viscosity    float64
	Turbulence   float64
	Repulsion    float64
	Speed        float64
	Size         float64
}

// Particle struct
type Particle struct {
	x, y   float64 // position
	vx, vy float64 // velocity
	color  color.RGBA
}

var particles []Particle

// Update runs the game logic
func (g *Game) Update() error {

	if time.Since(startTime).Milliseconds() > 0 {
		fmt.Println("FPS: ", 1000/float64(time.Since(startTime).Milliseconds()))
		startTime = time.Now()

	}

	// Mouse input
	mouseX, mouseY := ebiten.CursorPosition()
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		// Apply force to each particle toward the mouse position
		for i := range particles {
			// Calculate the direction vector from the particle to the mouse position
			dx := float64(mouseX) - particles[i].x
			dy := float64(mouseY) - particles[i].y

			// Calculate the distance (magnitude) of the vector
			distance := math.Hypot(dx, dy)

			// Avoid division by zero if particles are already at the mouse position
			if distance != 0 {
				// Normalize the direction vector and apply a force toward the mouse
				// We can scale the force based on distance to make particles move more slowly if they're farther away
				forceMagnitude := 1.0 / (distance + 1) // The "+1" prevents division by zero and adds some smoothing
				particles[i].vx += forceMagnitude * dx
				particles[i].vy += forceMagnitude * dy
			}
		}
	}
	// Apply viscosity (friction)
	for i := range particles {

		particles[i].vx *= conf.Viscosity
		particles[i].vy *= conf.Viscosity

		particles[i].vx += ((rand.Float64() - 0.5) * conf.Turbulence)
		particles[i].vy += ((rand.Float64() - 0.5) * conf.Turbulence)
		// Move particle
		particles[i].x += particles[i].vx
		particles[i].y += particles[i].vy

		// Edge detection
		if particles[i].x-conf.Size <= 0 || particles[i].x+conf.Size > float64(conf.Width)-conf.Size {
			particles[i].vx = -particles[i].vx
			// particles[i].x = float64(conf.Width / 2)
			// particles[i].y = float64(conf.Height / 2)
		}
		if particles[i].y-conf.Size <= 0 || particles[i].y+conf.Size > float64(conf.Height)-conf.Size {
			particles[i].vy = -particles[i].vy
			// particles[i].x = float64(conf.Width / 2)
			// particles[i].y = float64(conf.Height / 2)
		}

		// Repulsion between particles
		for j := range particles {
			if i == j {
				continue
			}
			distance := math.Hypot(particles[i].x-particles[j].x, particles[i].y-particles[j].y)
			if distance < 2*float64(conf.Size) {
				dx := particles[i].x - particles[j].x
				dy := particles[i].y - particles[j].y
				dist := math.Hypot(dx, dy)

				if dist != 0 {
					// Calculate force magnitude based on distance
					forceMagnitude := conf.Repulsion * (2*conf.Size - distance)
					// Apply force
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

// Draw renders the frame
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Clear()
	for _, p := range particles {
		if conf.DynamicColor {
			r := uint8(rand.Intn(256))
			g := uint8(rand.Intn(256))
			b := uint8(rand.Intn(256))
			vector.DrawFilledCircle(screen, float32(p.x), float32(p.y), float32(conf.Size), color.RGBA{r, g, b, 255}, true)
		} else {
			r := uint8(255)
			g := uint8(255)
			b := uint8(255)
			vector.DrawFilledCircle(screen, float32(p.x), float32(p.y), float32(conf.Size), color.RGBA{r, g, b, 127}, true)
		}
	}
}

// Layout sets screen size
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return conf.Width, conf.Height
}
