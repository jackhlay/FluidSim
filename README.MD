# Fluid Simulation  
  
#### A quick fluid simulation made in go. not quite water.  
  
In the "sim.go" file, there's a Config object with 12 params that will alter the simulation.  
 
## Config Params 
- DynamicColor(bool):  
    - Toggles dynamic color generation for each particle
- Width (int):  
    - Changes the width of the window
- Height (int):  
    - Changes the Height of the window
- Particles (int):  
    - Changes the Number of generated particles
- Viscosity (float64):  
    - A damping factor applied to particle velocities each frame (0–1, where 1 means no damping).  
    - Controls how quickly particle velocities decay over time. Higher values cause particles to "settle" slower.
- Turbulence (float64):  
    - Introduces random noise, to mimic movement in fluid systems.  
- Repulsion(float64):  
    - Defines the repulsive force strength between particles to keep them apart.
    - Higher values increase the distance particles attempt to maintain from each other
- Bounce (float64):  
    - Controls the elasticity of particles when they collide with boundaries.
    - A value of 1 means full elasticity (no energy loss), while values closer to 0 cause particles to lose more energy with each boundary collision. 
- Gravity (float64):  
    - The downward acceleration applied to each particle per frame.  
- SpringConstant (float64):  
    - The strength of the spring force applied between neighboring particles.  
    - Controls how strongly particles pull toward or push away from each other to reach an ideal distance (RestLength).
- RestLength (float64):  
    - The target distance between particles for the spring force.
    - If particles are closer than this distance, they push each other away; if farther, they pull toward each other.
- Size (float64):  
    - The radius of each particle in pixels.

## How parameters affect the simulation
- Cohesion and Structure:
    - ```SpringConstant``` and ```RestLength``` work together to determine the overall "structural integrity" of the fluid. Higher values for these parameters create a more solid, jelly-like substance, while lower values make the fluid looser.  
- Damping and Stability:
    - ```Viscosity``` and ```Repulsion``` influence the stability and smoothness of the particle motion. Increasing ```Viscosity``` makes particles settle faster and reduces erratic movement.  
  
## Controls
- Left Click  
    Attracts all particles to the mouse cursor
- Right Click  
    Repells all particles from the mouse cursor  

## To run
- You must have GoLang 1.20+ installed  
- Enter project repository
- Enter the following into the console:
```sh
go mod tidy
go run sim.go screen.go  
```