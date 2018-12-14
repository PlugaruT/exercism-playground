package strain

type Ints []int
type Lists [][]int
type Strings []string

func (ints Ints) Keep(f func(int) bool) (o Ints) {
	for _, number := range ints {
		if f(number) {
			o = append(o, number)
		}
	}
	return
}
func (lists Lists) Keep(f func([]int) bool) (o Lists) {
	for _, list := range lists {
		if f(list) {
			o = append(o, list)
		}
	}

	return
}
func (strings Strings) Keep(f func(string) bool) (o Strings) {
	for _, astring := range strings {
		if f(astring) {
			o = append(o, astring)
		}
	}

	return
}
func (ints Ints) Discard(f func(int) bool) (o Ints) {
	for _, number := range ints {
		if !f(number) {
			o = append(o, number)
		}
	}

	return
}
