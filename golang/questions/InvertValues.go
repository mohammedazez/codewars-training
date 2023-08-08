package questions

import (
	"math"
)

func Invert(arr []int) []int {
	var res []int
	for _, v := range arr {
		if v == v*-1 {
			res = append(res, int(math.Abs(float64(v))))
		} else {
			res = append(res, v*-1)
		}
	}
	return res
}
