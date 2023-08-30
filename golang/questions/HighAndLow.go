package questions

import (
	"sort"
	"strconv"
	"strings"
)

func HighAndLow(in string) string {
	split := strings.Split(in, " ")
	var temp []int
	for i := 0; i < len(split); i++ {
		conv, _ := strconv.Atoi(split[i])
		temp = append(temp, conv)
	}
	sort.Ints(temp)
	return strconv.Itoa(temp[len(temp)-1]) + " " + strconv.Itoa(temp[0])
}

// "1 2 -3 4 5"
// ["1", "2", "-3", "4", "5"]
// [1, 2, -3, 4, 5]
// [-3, 1, 2 , 4 , 5]
// 5, -3
