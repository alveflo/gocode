package counter

import "testing"

func TestCounter(t *testing.T) {
	str := "HELLO"

	var c Counter
	pairs := c.Count(str)
	if len(pairs) != 4 {
		t.Error("Expected length of 4!")
	}
	for _, p := range pairs {
		switch p.Character {
		case "H":
			if p.Repetitions != 1 {
				t.Error("Expected H to have repetition of 1")
			}
		case "E":
			if p.Repetitions != 1 {
				t.Error("Expected E to have repetition of 1")
			}
		case "L":
			if p.Repetitions != 2 {
				t.Error("Expected L to have repetition of 2")
			}
		case "O":
			if p.Repetitions != 1 {
				t.Error("Expected O to have repetition of 1")
			}
		default:
			t.Error("Unrecognized character!")
		}
	}
}