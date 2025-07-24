package main

import (
	"testing"
)

func BenchmarkMixedDataAccess(b *testing.B) {
	data := populateMixedData(N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = SumValuesMixedData(data)
	}
}

func BenchmarkOptimizedDataAccess(b *testing.B) {
	data := populateOptimizedData(N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = SumValuesOptimizedData(data)
	}
}
