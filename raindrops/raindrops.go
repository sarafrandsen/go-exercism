package raindrops

import "strconv"

// Convert takes an int and returns a string, printing "Pling", "Plang", or "Plong" if its factors include 3, 5, or 7
func Convert(n int) (s string) {
	s = ""
	if n%3 == 0 {
		s += "Pling"
	}
	if n%5 == 0 {
		s += "Plang"
	}
	if n%7 == 0 {
		s += "Plong"
	}
	if s == "" {
		s = strconv.Itoa(n)
	}
	return s
}
