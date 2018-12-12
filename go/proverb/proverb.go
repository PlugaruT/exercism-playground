package proverb

import (
	"fmt"
)

func formProverb(first, second string) string {
	var proverb string
	if second == "" {
		proverb = fmt.Sprintf("And all for the want of a %s.", first)
	} else {
		proverb = fmt.Sprintf("For want of a %s the %s was lost.", first, second)
	}
	return proverb
}

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
	if len(rhyme) == 0 {
		return []string{}
	}
	var response = []string{}

	for i := range rhyme {
		var curr = rhyme[i]
		var next = ""
		if !(len(rhyme)-1 == i) {
			next = rhyme[i+1]
		} else {
			curr = rhyme[0]
		}
		response = append(response, formProverb(curr, next))
	}

	return response
}
