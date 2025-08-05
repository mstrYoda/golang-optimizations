# Branch Prediction Optimization

Concept: Structure code to minimize unpredictable branches.

Scenario: Processing a slice of numbers with a condition that is either highly predictable or randomized.


## Bench Results

```bash
(base) emresavci@emre-mbp branch-prediction % go test -bench=. -benchmem .
goos: darwin
goarch: arm64
pkg: optimizations/branch-prediction
cpu: Apple M2
BenchmarkPredictableBranch-8       38072             28969 ns/op               0 B/op          0 allocs/op
BenchmarkRandomBranch-8            41528             28714 ns/op               0 B/op          0 allocs/op
BenchmarkBranchless-8              41641             29276 ns/op               0 B/op          0 allocs/op
PASS
ok      optimizations/branch-prediction 4.635s
```