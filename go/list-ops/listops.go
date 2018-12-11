package listops

// IntList represents a list containing only ints
type IntList []int

type binFunc func(x, y int) int
type predFunc func(x int) bool
type unaryFunc func(x int) int

// Length computes the len of the IntList
func (l IntList) Length() int {
	return len(l)
}

// Concat returns a new list with elements from l and other
func (l IntList) Concat(other []IntList) IntList {
	f := make(IntList, 0, len(l)*len(other))
	f = append(f, l...)
	for _, m := range other {
		f = append(f, m...)
	}
	return f
}

// Foldl reduces an integer list, starting with initial value acc, by
// applying binary function f repeatedly from left to right.
func (l IntList) Foldl(f binFunc, acc int) int {
	for _, i := range l {
		acc = f(acc, i)
	}
	return acc
}

// Foldr reduces an integer list, starting with initial value acc, by
// applying binary function f repeatedly from right to left.
func (l IntList) Foldr(f binFunc, acc int) int {
	for i := len(l) - 1; i >= 0; i-- {
		acc = f(l[i], acc)
	}
	return acc
}

// Map returns a list of the results applying function f for each element from l
func (l IntList) Map(f unaryFunc) IntList {
	result := make(IntList, 0, len(l))
	for _, val := range l {
		result = append(result, f(val))
	}
	return result
}

// Filter returns a list with elements from l for which the function f returns true
func (l IntList) Filter(f predFunc) IntList {
	result := make(IntList, 0, len(l))
	for _, val := range l {
		if f(val) {
			result = append(result, val)
		}
	}
	return result
}

// Reverse returns a new reversed list
func (l IntList) Reverse() IntList {
	result := make(IntList, len(l))
	for i, v := range l {
		result[len(l)-1-i] = v
	}
	return result
}

// Append returns a copy of l with the elements of other appended
func (l IntList) Append(other IntList) IntList {
	f := make(IntList, 0, len(l)+len(other))
	f = append(f, l...)
	f = append(f, other...)
	return f
}
