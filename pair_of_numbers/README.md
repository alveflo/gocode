# Pair of numbers
Finds matching pair from a given list of positive integers that adds up to a given sum
## Install
```
$ go get github.com/alveflo/gocode/pairs
```

## Test
```
$ cd pairs
$ go test
```

## Usage
```go
package main

import (
    "./pairs"
)

func main() {
    values := []int {10, 3, 5, 2, 4, 1, 9, 5, 6, 7, 0}
    sum := 10
    var finder PairFinder
    pairs := finder.FindPairs(values, sum)
}
```
