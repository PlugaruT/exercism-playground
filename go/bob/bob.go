// Package bob contains all the functions to simulate Bob the teenager
package bob

import "strings"

func sanitize(str string) string {
	return strings.TrimSpace(str)
}

func isUpper(str string) bool {
	return str == strings.ToUpper(str) && str != strings.ToLower(str)
}

func isBlank(str string) bool {
	return str == ""
}

func isQuestion(str string) bool {
	return strings.HasSuffix(str, "?")
}

func isYelledSentence(str string) bool {
	return isUpper(str) && !isQuestion(str)
}

func isYelledQuestion(str string) bool {
	return isUpper(str) && isQuestion(str)
}

// Hey simulates Bob, the lackadaisical teenager
func Hey(remark string) string {
	remark = sanitize(remark)
	switch {
	case isYelledSentence(remark):
		return "Whoa, chill out!"
	case isYelledQuestion(remark):
		return "Calm down, I know what I'm doing!"
	case isQuestion(remark):
		return "Sure."
	case isBlank(remark):
		return "Fine. Be that way!"
	}
	return "Whatever."
}
