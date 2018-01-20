package counter

type Counter struct {
}

type Pair struct {
	Character string
	Repetitions int
}

func (c Counter) Count(str string) [] Pair {
	var pairs []Pair
	repetitions := make([]int, 128)

	for i, _ := range str {
		repetitions[str[i]]++
	}

	for i, elem := range repetitions {
		if elem > 0 {
			pairs = append(pairs, Pair { string(i), elem })
		}
	}

	return pairs
}