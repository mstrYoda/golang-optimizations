package branchprediction

import "testing"

/*
To run benchmarks:
go test -bench=. -benchmem .
*/

func BenchmarkPredictableBranch(b *testing.B) {
	data := populatePredictableData(Size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = processPredictable(data)
	}
}

func BenchmarkRandomBranch(b *testing.B) {
	data := populateRandomData(Size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = processRandom(data)
	}
}

func BenchmarkBranchless(b *testing.B) {
	// For branchless, the input predictability doesn't matter much for the branch itself.
	// We'll use random data to compare against the worst-case for branching.
	data := populateRandomData(Size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = processBranchless(data)
	}
}
