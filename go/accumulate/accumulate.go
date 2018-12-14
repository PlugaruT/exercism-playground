package accumulate

// Accumulate receive as input a list of items and a func
// returns a new slice with items resulted of applying the func to each item
func Accumulate(items []string, f func(string) string) []string {
	var result = make([]string, 0, len(items))

	for _, item := range items {
		result = append(result, f(item))
	}
	return result
}
