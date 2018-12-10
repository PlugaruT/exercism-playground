package hamming

import "errors"

// Distance receives two strings and computes the hamming distance
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("The string length should be the same")
	}

	distance := 0
	for i := range a {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil
}
