package main

import (
	"fmt"
)

// Data structure with mixed data types
type MixedData struct {
	ID        int
	Value1    float64
	Name      string
	Timestamp int64
	Value2    float64
}

// Data structure optimized for cache locality (grouping frequently accessed types)
type OptimizedData struct {
	Value1 float64
	Value2 float64
	ID     int
	// Name and Timestamp might be less frequently accessed together,
	// or could be moved to another "hot" struct if their combination is frequent.
	Name      string
	Timestamp int64
}

const N = 100000 // Number of structs in our slice

// populateData creates a slice of MixedData
func populateMixedData(n int) []MixedData {
	data := make([]MixedData, n)
	for i := 0; i < n; i++ {
		data[i] = MixedData{
			ID:        i,
			Value1:    float64(i) * 1.1,
			Name:      fmt.Sprintf("Name-%d", i),
			Timestamp: int64(i * 100),
			Value2:    float64(i) * 2.2,
		}
	}
	return data
}

// populateOptimizedData creates a slice of OptimizedData
func populateOptimizedData(n int) []OptimizedData {
	data := make([]OptimizedData, n)
	for i := 0; i < n; i++ {
		data[i] = OptimizedData{
			Value1:    float64(i) * 1.1,
			Value2:    float64(i) * 2.2,
			ID:        i,
			Name:      fmt.Sprintf("Name-%d", i),
			Timestamp: int64(i * 100),
		}
	}
	return data
}

// SumValuesMixedData calculates sum of Value1 and Value2 from MixedData
// Accesses are spread out, potentially hitting different cache lines.
func SumValuesMixedData(data []MixedData) float64 {
	var sum float64
	for i := range data {
		sum += data[i].Value1 + data[i].Value2
	}
	return sum
}

// SumValuesOptimizedData calculates sum of Value1 and Value2 from OptimizedData
// Value1 and Value2 are contiguous in memory, improving cache locality.
func SumValuesOptimizedData(data []OptimizedData) float64 {
	var sum float64
	for i := range data {
		sum += data[i].Value1 + data[i].Value2
	}
	return sum
}
