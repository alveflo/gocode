package pairs

type PairFinder struct {
}

type Pair struct {
	A int
	B int
}

func (p PairFinder) FindPairs(list []int, sum int) []Pair {
	complements := make([]bool, sum + 1)
	var pairs []Pair

	for _, val := range list {
		if val < 0 {
			panic("No negative numbers are allowed")
		}

		if val <= sum {
			complement := sum - val

			if complements[val] {
				pairs = append(pairs, Pair {val, complement})
			}

			complements[complement] = true
		}
	}

	return pairs
}