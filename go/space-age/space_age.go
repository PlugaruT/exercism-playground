// Package space helps with calculating age in the age of space travel
package space

// Planet wraps a planet name
type Planet string

func orbitalPeriodInEarthYears(planet Planet) float64 {
	switch planet {
	case "Mercury":
		return 0.2408467
	case "Venus":
		return 0.61519726
	case "Mars":
		return 1.8808158
	case "Jupiter":
		return 11.862615
	case "Saturn":
		return 29.447498
	case "Uranus":
		return 84.016846
	case "Neptune":
		return 164.79132
	}
	return 1
}

// Age returns someone's age in years on a different planet
func Age(seconds float64, planet Planet) float64 {
	return seconds / (31557600 * orbitalPeriodInEarthYears(planet))
}
