package main

import (
	"fmt"
	"./pairs"
)

func main() {
	values := []int { 2, 8, 11, 3, 2, 5, 6, 0 }
	sum := 8
	var finder pairs.PairFinder
	pairs := finder.FindPairs(values, sum)

	fmt.Println("Given")
	fmt.Print("  [ ")
	for _, val := range values {
		fmt.Print(val, " ")
	}
	fmt.Println("]\n\nLooking for x + y = ", sum)

	if len(pairs) > 0 {		
		fmt.Println("\n- Matching pairs:")
		for _, p := range pairs {
			fmt.Println("\t", p.A, "+", p.B, "=", sum)
		}
	} else {
		fmt.Println("\n- No pair was found")
	}
}