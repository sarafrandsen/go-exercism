package gigasecond

import (
	"math"
	"time"
)

// AddGigasecond calculates the moment when someone has lived for 10^9 seconds.
func AddGigasecond(t time.Time) time.Time {
	g := time.Duration(math.Pow(10, 9))
	t = t.Add((time.Second) * g)
	return t
}
