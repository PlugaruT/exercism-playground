package raindrops

import "strconv"

const testVersion = 3

func isDivBy(nr int, divBy int) bool {
	return nr%divBy == 0
}

// Convert a number to a string, the contents of which depend on the number's factors.
func Convert(i int) string {
	var result string
	if isDivBy(i, 3) {
		result += "Pling"
	}
	if isDivBy(i, 5) {
		result += "Plang"
	}
	if isDivBy(i, 7) {
		result += "Plong"
	}
	if len(result) == 0 {
		return strconv.Itoa(i)
	}
	return result
}
