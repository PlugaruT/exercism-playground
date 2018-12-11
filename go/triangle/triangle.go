package triangle

import "math"

// Kind represents the type of the triangle
type Kind int

// Types of triangles
const (
	NaT = iota // not a triangle
	Equ        // equilateral
	Iso        // isosceles
	Sca        // scalene
)

// isNaT determines if it isn't a triangle
func isNaT(a, b, c float64) bool {
	var someNaN = math.IsNaN(a) ||
		math.IsNaN(b) ||
		math.IsNaN(c)
	var someNegative = (a <= 0 && b <= 0 && c <= 0)
	var someInf = math.IsInf(a, 0) ||
		math.IsInf(b, 0) ||
		math.IsInf(c, 0)
	var notATriangle = a+b < c || b+c < a || c+a < b
	return someNaN || someNegative || someInf || notATriangle
}

// isEqu determines is the triangle is Equilateral
func isEqu(a, b, c float64) bool {
	return a == b && b == c
}

// isIso determines is a triangle is Isosceles
func isIso(a, b, c float64) bool {
	return a == b ||
		b == c ||
		c == a
}

// isSca determines if a tringale is Scalene
func isSca(a, b, c float64) bool {
	return a != b &&
		b != c &&
		c != a
}

// KindFromSides checks the type of the triange based on a, b, c
func KindFromSides(a, b, c float64) Kind {
	var k Kind
	if isIso(a, b, c) {
		k = Iso
	}
	if isEqu(a, b, c) {
		k = Equ
	}
	if isSca(a, b, c) {
		k = Sca
	}
	if isNaT(a, b, c) {
		k = NaT
	}
	return k
}
