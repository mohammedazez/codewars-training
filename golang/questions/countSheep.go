package questions

import (
	"strconv"
)

func CountSheep(num int) string {
	// Your code here!
	var result string
	for i := 1; i <= num; i++ {
		result += strconv.Itoa(i) + " sheep..."
	}
	return result
}
