package dataorienteddesign

const EntityCount = 100000

// --- Traditional OOP-like approach (Array of Structs - AoS) ---

type Position struct {
	X, Y, Z float64
}

type Velocity struct {
	DX, DY, DZ float64
}

type Health struct {
	Current, Max float64
}

type GameEntity struct {
	ID       int
	Position Position
	Velocity Velocity
	Health   Health
	// Other fields...
}

// UpdatePositionsAoS iterates through a slice of GameEntity to update positions.
// This accesses disparate parts of memory for each entity (Position, Velocity).
func UpdatePositionsAoS(entities []GameEntity, deltaTime float64) {
	for i := range entities {
		entities[i].Position.X += entities[i].Velocity.DX * deltaTime
		entities[i].Position.Y += entities[i].Velocity.DY * deltaTime
		entities[i].Position.Z += entities[i].Velocity.DZ * deltaTime
	}
}

// --- Data-Oriented Design (Array of Component Slices - AoCS) ---

type Positions []struct{ X, Y, Z float64 }
type Velocities []struct{ DX, DY, DZ float64 }
type Healths []struct{ Current, Max float64 }

// UpdatePositionsAoCS iterates through separate slices for Position and Velocity.
// This accesses contiguous memory blocks for X, Y, Z, DX, DY, DZ respectively,
// which is much more cache-friendly when only these components are needed.
func UpdatePositionsAoCS(positions Positions, velocities Velocities, deltaTime float64) {
	for i := range positions {
		positions[i].X += velocities[i].DX * deltaTime
		positions[i].Y += velocities[i].DY * deltaTime
		positions[i].Z += velocities[i].DZ * deltaTime
	}
}

// populateAoS creates a slice of GameEntity
func populateAoS(count int) []GameEntity {
	entities := make([]GameEntity, count)
	for i := 0; i < count; i++ {
		entities[i] = GameEntity{
			ID: i,
			Position: Position{
				X: float64(i), Y: float64(i), Z: float64(i)},
			Velocity: Velocity{
				DX: float64(i%10 + 1), DY: float64(i%10 + 1), DZ: float64(i%10 + 1)},
			Health: Health{Current: 100, Max: 100},
		}
	}
	return entities
}

// populateAoCS creates separate slices for components
func populateAoCS(count int) (Positions, Velocities, Healths) {
	positions := make(Positions, count)
	velocities := make(Velocities, count)
	healths := make(Healths, count)

	for i := 0; i < count; i++ {
		positions[i] = struct{ X, Y, Z float64 }{X: float64(i), Y: float64(i), Z: float64(i)}
		velocities[i] = struct{ DX, DY, DZ float64 }{DX: float64(i%10 + 1), DY: float64(i%10 + 1), DZ: float64(i%10 + 1)}
		healths[i] = struct{ Current, Max float64 }{Current: 100, Max: 100}
	}
	return positions, velocities, healths
}
