package questions

import "sort"

func ExpressionMatter(a int, b int, c int) int {
	// your code here
	combinations := []int{
		a + b + c,
		a * b * c,
		(a + b) * c,
		a * (b + c),
		a*b + c,
		a + b*c,
	}

	// Find the maximum value from the combinations
	maxValue := combinations[0]
	for _, val := range combinations {
		if val > maxValue {
			maxValue = val
		}
	}
	return maxValue
}

// simple case
func ExpressionMatter2(a int, b int, c int) int {
	arr := []int{a * (b + c), a * b * c, a + b + c, (a + b) * c}
	sort.Ints(arr)
	return arr[len(arr)-1]
}
