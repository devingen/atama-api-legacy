package util

import "math"

func MaxIterationLimit(m int) int {
	v := 20 * math.Sqrt(float64(m)) / float64(m)
	return int(math.Round(v))
}

func MaxIterationLevel(m int) int {
	v := 60 * math.Sqrt(float64(m)) / float64(m)
	return int(math.Round(v))
}
