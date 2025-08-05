package branchprediction

import (
	"math/rand"
	"time"
)

const Size = 100000 // Size of the array

// processPredictable processes numbers where condition is highly predictable (e.g., all even).
// This allows the CPU's branch predictor to work effectively.
func processPredictable(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		if n%2 == 0 { // Highly predictable branch (always true if numbers are all even)
			sum += n
		} else {
			sum -= n // This path is rarely taken
		}
	}
	return sum
}

// processRandom processes numbers where condition is random (e.g., half even, half odd).
// This makes branch prediction difficult for the CPU.
func processRandom(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		if n%2 == 0 { // Unpredictable branch
			sum += n
		} else {
			sum -= n
		}
	}
	return sum
}

// processBranchless processes numbers using bitwise operations or other branchless techniques.
// This avoids the branch predictor entirely.
func processBranchless(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		// (n%2 == 0) is equivalent to (n&1 == 0) for non-negative integers
		// For branchless: if n is even, add n. If odd, subtract n.
		// Can be done with a conditional move or arithmetic trick:
		// If n is even, (n&1) is 0. If n is odd, (n&1) is 1.
		// (1 - 2*(n&1)) gives 1 for even, -1 for odd.
		// sum += n * (1 - 2*(n&1)) is a branchless way to achieve this.
		sum += n * (1 - 2*(n&1))
	}
	return sum
}

// populatePredictableData creates a slice of all even numbers
func populatePredictableData(size int) []int {
	data := make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = i * 2 // All even
	}
	return data
}

// populateRandomData creates a slice with roughly half even, half odd numbers
func populateRandomData(size int) []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	data := make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = r.Intn(size) // Random numbers
	}
	return data
}
