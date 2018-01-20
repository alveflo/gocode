package main

import (
	"fmt"
	"./counter"
)

func main() {
	str := "Hello world!"
	var c counter.Counter
	pairs := c.Count(str)

	fmt.Println("Given", str)

	fmt.Println("\n- Repetition per character:")
	for _, p := range pairs {
		fmt.Println("\t", p.Character, ":", p.Repetitions)
	}
}