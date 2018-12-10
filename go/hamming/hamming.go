package hamming

import "errors"

// Distance receives two strings and computes the hamming distance
func Distance(a, b string) (int, error) {
	distance := 0
	if len(a) != len(b) {
		return 0, errors.New("The string length should be the same")
	}

	for index := 0; index < len(a); index++ {
		if a[index] != b[index] {
			distance++
		}
	}
	return distance, nil
}
