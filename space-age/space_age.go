package space

// Planet represents a planet
type Planet string

const secondsPerEarthYear = 31557600.0

func round(seconds float64) float64 {
	return float64(int64(seconds/0.01+0.5)) * 0.01
}

// Age compares time in a planet's orbit to Earth's orbit
func Age(seconds float64, planet Planet) float64 {
	planets := map[Planet]float64{
		"Earth":   1.0,
		"Mercury": 0.2408467,
		"Venus":   0.61519726,
		"Mars":    1.8808158,
		"Jupiter": 11.862615,
		"Saturn":  29.447498,
		"Uranus":  84.016846,
		"Neptune": 164.79132,
	}
	return round(seconds / (secondsPerEarthYear * planets[planet]))
}
