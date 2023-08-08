package questions

import (
	"math"
)

func PowersOfTwo(n int) []uint64 {
	// your code goes here
	var res []uint64
	for i := 0; i <= n; i++ {
		res = append(res, uint64(math.Pow(2, float64(i))))
	}
	return res
}
