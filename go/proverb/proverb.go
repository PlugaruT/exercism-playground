package proverb

import (
	"fmt"
)

func formProverb(first, second string) string {
	shortVersion := "And all for the want of a %s."
	logVersion := "For want of a %s the %s was lost."
	proverb := fmt.Sprintf(shortVersion, first)
	if second != "" {
		proverb = fmt.Sprintf(logVersion, first, second)
	}
	return proverb
}

func nextOrEmpty(arr []string, index int) string {
	var next = ""
	if !(len(arr)-1 == index) {
		next = arr[index+1]
	}
	return next
}

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
	var response = []string{}
	for i := range rhyme {
		var curr = rhyme[i]
		var next = nextOrEmpty(rhyme, i)
		if len(rhyme)-1 == i {
			curr = rhyme[0]
		}
		response = append(response, formProverb(curr, next))
	}
	return response
}
