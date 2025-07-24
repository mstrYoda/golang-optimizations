package dataorienteddesign

import "testing"

/*
To run benchmarks:
go test -bench=. -benchmem .
*/

func BenchmarkUpdatePositionsAoS(b *testing.B) {
	entities := populateAoS(EntityCount)
	deltaTime := 0.01
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		UpdatePositionsAoS(entities, deltaTime)
	}
}

func BenchmarkUpdatePositionsAoCS(b *testing.B) {
	positions, velocities, _ := populateAoCS(EntityCount)
	deltaTime := 0.01
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		UpdatePositionsAoCS(positions, velocities, deltaTime)
	}
}
