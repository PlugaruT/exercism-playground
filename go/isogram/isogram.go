package isogram

import "unicode"

// IsIsogram checks if a strings is an isogram
func IsIsogram(in string) bool {
	seen := make(map[rune]bool)
	for _, r := range in {
		if !unicode.IsLetter(r) {
			continue
		}
		if _, ok := seen[unicode.ToLower(r)]; ok {
			return false
		}
		seen[unicode.ToLower(r)] = true
	}
	return true
}
