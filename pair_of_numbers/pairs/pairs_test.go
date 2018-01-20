package pairs

import "testing"

func TestFindPairs(t *testing.T) {
	values := []int {10, 3, 5, 2, 4, 1, 9, 5, 6, 7, 0}
	sum := 10
	var finder PairFinder
	pairs := finder.FindPairs(values, sum)
	if len(pairs) != 5 {
		t.Error("Expected length of 5!")
	}
}