// space package provides API to calculate age according to each planet
package space

// Planet is a type alias of string
// this makes Age signature more readable
type Planet string

// Age calculates age according to each planet
func Age(seconds float64, planet Planet) float64 {
	weightedYears := map[Planet]float64{
		"Mercury": 0.2408467,
		"Venus":   0.61519726,
		"Earth":   1.0,
		"Mars":    1.8808158,
		"Jupiter": 11.862615,
		"Saturn":  29.447498,
		"Uranus":  84.016846,
		"Neptune": 164.79132,
	}
	if _, ok := weightedYears[planet]; !ok {
		return -1
	}
	return seconds / 31557600 / weightedYears[planet]
}
