package main

import "fmt"

type pair struct {
	a int
	b int
}

func findPairs(list []int, sum int) []pair {
	complements := make([]bool, sum + 1)
	var pairs []pair

	for _, val := range list {
		if val < 0 {
			panic("No negative numbers are allowed")
		}

		if val <= sum {
			complement := sum - val

			if complements[val] {
				pairs = append(pairs, pair {val, complement})
			}

			complements[complement] = true
		}
	}

	return pairs
}

func main() {
	values := []int { 2, 8, 11, 3, 2, 5, 6, 0 }
	sum := 8
	pairs := findPairs(values, sum)

	fmt.Println("Given")
	fmt.Print("[ ")
	for _, val := range values {
		fmt.Print(val, " ")
	}
	fmt.Println("]\n\nLooking for x + y = ", sum)

	if len(pairs) > 0 {		
		fmt.Println("\n- Matching pairs:")
		for _, p := range pairs {
			fmt.Println("\t", p.a, "+", p.b, "=", sum)
		}
	} else {
		fmt.Println("\n- No pair was found")
	}
}