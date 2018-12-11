package collatzconjecture

import "fmt"

// CollatzConjecture solves the 3x+1 problem
func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, fmt.Errorf("Value for %d should be greater than 0", n)
	}

	steps := 0

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
