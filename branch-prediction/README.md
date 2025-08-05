# Branch Prediction Optimization

Concept: Structure code to minimize unpredictable branches.

Scenario: Processing a slice of numbers with a condition that is either highly predictable or randomized.

# Optimization Tricks

*Profile first*: Don't guess. Use go tool pprof to find the hotspots. Only optimize if a function is a genuine bottleneck.

*Make branches predictable*: Order your if/else statements with the most likely case first.

*Eliminate branches*: Use branchless techniques (bitwise, arithmetic) or lookup tables for simple logic in hot loops.

*Be aware of dynamic dispatch*: If interface calls are a bottleneck, consider whether you can use a concrete type directly.

*Use PGO*: Let the compiler do the heavy lifting for you by providing it with real-world execution data.


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