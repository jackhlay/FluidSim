package main

import (
	"image/color"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

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
	Bounce       float64
	Gravity      float64
	Size         float64
}

// Particle struct
type Particle struct {
	x, y   float64 // position
	vx, vy float64 // velocity
	color  color.RGBA
}

var particles []Particle

// Spatial partitioning grid to reduce pairwise repulsion calculations
var gridWidth, gridHeight int
var grid [][]int // each cell contains indices of particles

// Initialize particles and grid
func initParticles() {
	gridWidth, gridHeight = int(float64(conf.Width)/conf.Size), int(float64(conf.Height)/conf.Size)
	grid = make([][]int, gridWidth*gridHeight)

	for i := 0; i < conf.Particles; i++ {
		p := Particle{
			x:     rand.Float64() * float64(conf.Width),
			y:     rand.Float64() * float64(conf.Height),
			vx:    (rand.Float64() - 0.5) * conf.Turbulence,
			vy:    (rand.Float64() - 0.5) * conf.Turbulence,
			color: getRandomColor(),
		}
		particles = append(particles, p)
	}
}

// Get a random color for particles with `DynamicColor` enabled
func getRandomColor() color.RGBA {
	if conf.DynamicColor {
		return color.RGBA{uint8(rand.Intn(256)), uint8(rand.Intn(256)), uint8(rand.Intn(256)), 255}
	}
	return color.RGBA{0, 242, 255, 120}
}

// Update runs the game logic
func (g *Game) Update() error {
	// Calculate and display FPS every second
	// if time.Since(startTime).Milliseconds() > 0 {
	// 	fps = float64(1000) / float64(time.Since(startTime).Milliseconds()+1)
	// 	fmt.Printf("FPS: %.2f\n", fps)
	// 	startTime = time.Now()
	// }

	// Reset the spatial partitioning grid
	for i := range grid {
		grid[i] = nil
	}

	// Update particle grid positions
	for i := range particles {
		p := &particles[i]
		gridIndex := int(p.x/conf.Size)*gridHeight + int(p.y/conf.Size)
		grid[gridIndex] = append(grid[gridIndex], i)
	}

	// Mouse input
	mouseX, mouseY := ebiten.CursorPosition()
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		for i := range particles {
			p := &particles[i]
			p.vy -= conf.Gravity

			// Apply force towards the mouse
			dx, dy := float64(mouseX)-p.x, float64(mouseY)-p.y
			distanceSquared := dx*dx + dy*dy

			if distanceSquared > 0 {
				forceMagnitude := 1.0 / math.Sqrt(distanceSquared+1)
				p.vx += forceMagnitude * dx
				p.vy += forceMagnitude * dy
			}
		}
	}

	// Apply gravity, viscosity, and turbulence, and move particles
	for i := range particles {
		p := &particles[i]

		p.vy += conf.Gravity
		p.vx *= conf.Viscosity
		p.vy *= conf.Viscosity
		p.vx += (rand.Float64() - 0.5) * conf.Turbulence
		p.vy += (rand.Float64() - 0.5) * conf.Turbulence

		p.x += p.vx
		p.y += p.vy

		// Edge detection and boundary handling
		if p.x-conf.Size <= 0 {
			p.vx = -p.vx
			p.x = conf.Size
		} else if p.x+conf.Size >= float64(conf.Width) {
			p.vx = -p.vx
			p.x = float64(conf.Width) - conf.Size
		}
		if p.y-conf.Size <= 0 {
			p.vy = -p.vy
			p.y = conf.Size
		} else if p.y+conf.Size >= float64(conf.Height) {
			p.vy = -p.vy * conf.Bounce
			p.y = float64(conf.Height) - conf.Size
		}
	}

	// Particle repulsion with spatial partitioning
	for i := range particles {
		p := &particles[i]
		gridIndex := int(p.x/conf.Size)*gridHeight + int(p.y/conf.Size)

		for _, j := range grid[gridIndex] {
			if i == j {
				continue
			}
			p2 := &particles[j]
			dx, dy := p.x-p2.x, p.y-p2.y
			distanceSquared := dx*dx + dy*dy

			if distanceSquared < 4*conf.Size*conf.Size && distanceSquared > 0 {
				forceMagnitude := conf.Repulsion / distanceSquared
				p.vx += forceMagnitude * dx
				p.vy += forceMagnitude * dy
				p2.vx -= forceMagnitude * dx
				p2.vy -= forceMagnitude * dy
			}
		}
	}

	return nil
}

// Draw renders the frame
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Clear()
	for _, p := range particles {
		vector.DrawFilledCircle(screen, float32(p.x), float32(p.y), float32(conf.Size), p.color, false)
	}
}

// Layout sets screen size
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return conf.Width, conf.Height
}
