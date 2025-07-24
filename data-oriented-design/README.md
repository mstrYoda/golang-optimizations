# Data-Oriented Design (DOD)

Concept: Organize data around how it will be processed, not just object relationships.

Scenario: Representing game entities with traditional OOP-like structs vs. separate slices for components (AoCS - Array of Component Slices).

## Benc Results

```bash
(base) emresavci@emre-mbp data-oriented-design % go test -bench=. -benchmem .
goos: darwin
goarch: arm64
pkg: optimizations/data-oriented-design
cpu: Apple M2
BenchmarkUpdatePositionsAoS-8               8322            141272 ns/op               0 B/op          0 allocs/op
BenchmarkUpdatePositionsAoCS-8              9298            123865 ns/op               0 B/op          0 allocs/op
PASS
ok      optimizations/data-oriented-design      3.764s
```