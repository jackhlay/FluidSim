# Fluid Simulation  
  
#### A quick simulation made in go. not too sopishtocated.  
  
In the "sim.go" file, there's a Config object with 8 params that will alter the simulation.  
 
## Config Params 
- DynamicColor(bool):  
    Toggles dynamic color generation for each particle
- Width (int):
    Changes the width of the window
- Height (int):
    Changes the Height of the window
- Particles (int):
    Changes the Number of generated particles
- Viscosity (float64):  
    Slows or speeds up movement
- Turbulence (float64):  
    Introduces random noise, to mimic movement in fluid systems.  
- Repulsion(float64):  
    Creates interactions between particles
- Size (float64):  
    Determines radius of particles  
  
## To run
- You must have GoLang installed  
- Enter project repository
- Enter the following into the console:
```sh
go mod tidy
go run sim.go screen.go  
```