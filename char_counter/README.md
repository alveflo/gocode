# Pair of numbers
Counts repetition of each characters in a given string

## Test
```
$ cd counter
$ go test
```

## Usage
```go
package main

import (
    "fmt"
    "./counter"
)

func main() {
	str := "Hello world!"
	var c counter.Counter
	pairs := c.Count(str)
}
```
