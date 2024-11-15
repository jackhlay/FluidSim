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

	lowVelocityThreshold float64 // When particles fall below this speed, apply friction

	Friction  float64 // Friction coefficient to slow down particles
	Drag      float64
	Repulsion float64
	Bounce    float64
	Gravity   float64

	SpringConstant float64 // Cohesion force constant
	RestLength     float64 // Ideal distance between particles for stability

	Size float64
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

// Dynamically change particle color based on speed
var lowSpeedColor = color.RGBA{0, 57, 255, 255}
var highSpeedColor = color.RGBA{255, 0, 0, 255}

func interpColors(c1, c2 color.RGBA, t float64) color.RGBA {
	return color.RGBA{
		uint8(float64(c1.R)*(1-t) + float64(c2.R)*t),
		uint8(float64(c1.G)*(1-t) + float64(c2.G)*t),
		uint8(float64(c1.B)*(1-t) + float64(c2.B)*t),
		255,
	}
}

// Initialize particles and grid
func initParticles() {
	gridWidth, gridHeight = int(float64(conf.Width)/conf.Size)+1, int(float64(conf.Height)/conf.Size)+1
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
	return color.RGBA{0, 49, 144, 190}
}

// Get grid cell index from particle coordinates
func getGridCell(x, y float64) int {
	col := int(x / conf.Size)
	row := int(y / conf.Size)
	if col >= gridWidth {
		col = gridWidth - 1
	} else if col < 0 {
		col = 0
	}
	if row >= gridHeight {
		row = gridHeight - 1
	} else if row < 0 {
		row = 0
	}
	return col*gridHeight + row
}

// Get indices of neighboring particles within adjacent cells
func getNeighbors(gridCell int) []int {
	neighbors := []int{}

	col := gridCell / gridHeight
	row := gridCell % gridHeight

	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			neighborCol := col + dx
			neighborRow := row + dy

			// Ensure the neighbor cell is within grid bounds
			if neighborCol >= 0 && neighborCol < gridWidth && neighborRow >= 0 && neighborRow < gridHeight {
				neighborCell := neighborCol*gridHeight + neighborRow
				neighbors = append(neighbors, grid[neighborCell]...)
			}
		}
	}
	return neighbors
}

// Update runs the game logic
func (g *Game) Update() error {
	maxSpeed := 0.0
	for i := range grid {
		grid[i] = nil
	}

	// Update particle grid positions
	for i := range particles {
		p := &particles[i]
		speed := math.Hypot(p.vx, p.vy)
		if speed > maxSpeed {
			maxSpeed = speed
		}
		p.vx *= (1 - conf.Drag)
		p.vy *= (1 - conf.Drag)

	}
	for i := range particles {
		p := &particles[i]
		if conf.DynamicColor {
			speed := math.Hypot(p.vx, p.vy)
			normalizedSpeed := speed / maxSpeed
			particles[i].color = interpColors(lowSpeedColor, highSpeedColor, normalizedSpeed)
		}

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
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		for i := range particles {
			p := &particles[i]
			p.vy -= conf.Gravity

			// Apply force aWay from the mouse
			dx, dy := float64(mouseX)-p.x, float64(mouseY)-p.y
			distanceSquared := dx*dx + dy*dy

			if distanceSquared > 0 {
				forceMagnitude := 1.0 / math.Sqrt(distanceSquared+1)
				p.vx += forceMagnitude * -2 * dx
				p.vy += forceMagnitude * -2 * dy
			}
		}
	}

	// Apply gravity, viscosity, and turbulence, and move particles
	for i := range particles {
		p := &particles[i]

		p.vy += conf.Gravity
		p.vx *= conf.Viscosity
		p.vy *= conf.Viscosity

		// Apply drag
		// p.vx += (rand.Float64() - 0.5)
		// p.vy += (rand.Float64() - 0.5)
		// Apply spring force to neighbors within a small distance
		gridCell := getGridCell(particles[i].x, particles[i].y)
		for _, neighbor := range getNeighbors(gridCell) {
			if i == neighbor {
				continue
			}
			dx := particles[neighbor].x - particles[i].x
			dy := particles[neighbor].y - particles[i].y

			distance := math.Hypot(dx, dy)
			dynamicSpringConstant := conf.SpringConstant * (distance / conf.RestLength)

			// Spring force
			if distance < conf.RestLength && distance > 0 {
				force := dynamicSpringConstant * (distance - conf.RestLength)
				fx := (dx / distance) * force
				fy := (dy / distance) * force

				particles[i].vx += fx
				particles[i].vy += fy
				particles[neighbor].vx -= fx
				particles[neighbor].vy -= fy
			}
		}

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
			p.vy = -p.vy * conf.Bounce
			p.y = conf.Size
		} else if p.y+conf.Size >= float64(conf.Height) {
			p.vy = -p.vy
			p.y = float64(conf.Height) - conf.Size
		}
	}

	return nil
}

// Draw renders the frame
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Clear()
	for _, p := range particles {
		vector.DrawFilledCircle(screen, float32(p.x), float32(p.y), float32(conf.Size), p.color, true)
	}
}

// Layout sets screen size
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return conf.Width, conf.Height
}
