package questions

import (
	"math"
)

func MultipleOfIndex(ints []int) []int {
	// good luck
	var results []int

	for i, v := range ints {
		if i != 0 {
			var sum float64 = float64(v) / float64(i)
			if math.Floor(sum) == sum {
				results = append(results, int(v))
			}
		}
	}

	return results
}
