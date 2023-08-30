package questions

import (
	"math"
	"strings"
)

func GetMiddle(s string) string {
	var temp []float64
	var res string
	split := strings.Split(s, "")
	for i := 0; i < len(split); i++ {
		temp = append(temp, float64(i))
	}
	middle := math.Floor(float64(len(temp) / 2))

	if len(s)%2 == 0 {
		res = string(s[int(middle)-1]) + string(s[int(middle)])
	} else {
		res = string(s[int(middle)])
	}

	return res
}

// jika genap return ini
// "test"  -> es
// ["t", "e", "s", "t"]
// [0, 1, 2, 3]
// [1 , 2]
// 1 2 => es

// jika ganjil return ini
// "testing" -> t
// ["t", "e", "s", "t", "i", "n", "g"]
// [0, 1, 2, 3, 4, 5, 6]
// [3]
// 3 => t
