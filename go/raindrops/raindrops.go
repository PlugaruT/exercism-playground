package raindrops

import "strings"
import "strconv"

const testVersion = 3

func isDivBy(nr int, divBy int) bool {
	return nr%divBy == 0
}

// Convert a number to a string, the contents of which depend on the number's factors.
func Convert(i int) string {
	s := make([]string, 0)
	if isDivBy(i, 3) {
		s = append(s, "Pling")
	}
	if isDivBy(i, 5) {
		s = append(s, "Plang")
	}
	if isDivBy(i, 7) {
		s = append(s, "Plong")
	}
	if len(s) == 0 {
		return strconv.Itoa(i)
	}
	return strings.Join(s, "")
}
