// Package acronym helps with getting an acronym for any three-letter-string
package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate generates a three letter acronym of any name
func Abbreviate(s string) string {
	splitRegexp := regexp.MustCompile(`\W`)
	wordsSlice := splitRegexp.Split(s, -1)

	var response strings.Builder
	for _, word := range wordsSlice {
		if len(word) > 1 {
			response.WriteString(strings.ToUpper(word[0:1]))
		}
	}
	return response.String()
}
