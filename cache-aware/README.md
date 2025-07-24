# Cache Aware Programming

Concept: Design data structures and access patterns to maximize CPU cache hits.

Scenario: Iterating over a slice of structs where fields are accessed randomly vs. contiguously.

## Bench Results

```bash
(base) emresavci@emre-mbp cache-aware % go test -bench=. -benchmem .
goos: darwin
goarch: arm64
pkg: optimizations/cache-aware
cpu: Apple M2
BenchmarkMixedDataAccess-8                 41378             28791 ns/op               0 B/op          0 allocs/op
BenchmarkOptimizedDataAccess-8             41671             29972 ns/op               0 B/op          0 allocs/op
PASS
ok      optimizations/cache-aware       3.502s
```