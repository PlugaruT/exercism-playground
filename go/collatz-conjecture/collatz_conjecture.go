package collatzconjecture

import "errors"

// CollatzConjecture solves the 3x+1 problem
func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("Value for n should be greater than 0")
	}

	steps := 0

	if n == 1 {
		return steps, nil
	}

	for n > 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
		steps++
	}
	return steps, nil
}
