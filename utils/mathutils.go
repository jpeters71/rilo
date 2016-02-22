package utils

import "math"

// Round takes a float64 and rounds it to the nearest integer.  It's REALLY annoying this isn't part
// of the standard Go library.  Seriously, I don't see why they would reject particularly since all
// of Go's math functions take float64 values...it's not uncommon to want to do integer arithmetic.
func Round(f float64) int {
	if math.Abs(f) < 0.5 {
		return 0
	}
	return int(f + math.Copysign(0.5, f))
}
