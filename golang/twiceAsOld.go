package main

import "math"

func TwiceAsOld(dadYearsOld, sonYearsOld int) int {
	// Implement me
	result := dadYearsOld - (sonYearsOld)*2
	coba := math.Abs(float64(result))
	return int(coba)
}
