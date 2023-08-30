package questions

import (
	"math/bits"
)

func RangeBitCount(a, b int) (result int) {
	for i := a; i <= b; i++ {
		result += bits.OnesCount(uint(i))
	}
	return
}
