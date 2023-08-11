package questions

import (
	"fmt"
	"math"
)

func Psbb(families int, memberFamilies []int) {
	if families != len(memberFamilies) {
		fmt.Println("Input must be equal with count of family")
	} else {
		var res float64
		for _, v := range memberFamilies {
			res += float64(v)
		}

		fmt.Println("Minimum bus required is :", math.Ceil(res/4))
	}
}
