## simple blockchain

[![Go Report Card](https://goreportcard.com/badge/jancajthaml-go/blockchain)](https://goreportcard.com/report/jancajthaml-go/blockchain)

Simplest blockchain

### Usage ###

```
import "github.com/jancajthaml-go/blockchain"

ok := luhn.Validate("00123014764700968325")

digit, error := luhn.Digit("x")

checksum := luhn.Generate("1")
```

### Performance ###

```
BenchmarkLuhnSmall-4          200000000  13.9 ns/op  0 B/op  0 allocs/op
BenchmarkLuhnLarge-4          50000000   31.8 ns/op  0 B/op  0 allocs/op
BenchmarkLuhnSmallParallel-4  300000000  4.06 ns/op  0 B/op  0 allocs/op
BenchmarkLuhnLargeParallel-4  100000000  15.0 ns/op  0 B/op  0 allocs/op
```

verify your performance by running `make benchmark`

### Resources ###

* [Wikipedia - Blockchain](https://en.wikipedia.org/wiki/Blockchain)
