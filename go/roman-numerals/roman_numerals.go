package romannumerals

import "fmt"

// ToRomanNumeral converts a number from it's roman representation
func ToRomanNumeral(number int) (string, error) {
	conversions := []struct {
		value int
		digit string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var result string
	if number <= 0 || number > 3000 {
		return result, fmt.Errorf("number should be in the intervar 1 and 3000")
	}

	for _, conversion := range conversions {
		for number >= conversion.value {
			result += conversion.digit
			number -= conversion.value
		}
	}

	return result, nil
}
