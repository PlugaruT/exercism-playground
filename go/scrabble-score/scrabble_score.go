package scrabble

import "strings"

type point struct {
	letters string
	value   int
}

var scrabbleValues = []point{
	{"aeioulnrst", 1},
	{"dg", 2},
	{"bcmp", 3},
	{"fhvwy", 4},
	{"k", 5},
	{"jx", 8},
	{"qz", 10},
}

// Score computes the scrabble score for that word.
func Score(s string) int {
	var score int
	for _, r := range s {
		c := strings.ToLower(string(r))
		for _, rule := range scrabbleValues {
			if strings.ContainsAny(rule.letters, c) {
				score += rule.value
			}
		}
	}
	return score
}
