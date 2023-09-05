package arrayandhashing

import "fmt"

// corret answer
func ReplaceElements2(arr []int) []int {
	rightMax := -1
	for i := len(arr) - 1; i >= 0; i-- {
		newMax := rightMax
		if arr[i] > rightMax {
			newMax = arr[i]
		}
		arr[i] = rightMax
		rightMax = newMax
	}
	return arr
}

// [17, 18, 5, 4, 6, 1 ] -> [18, 6, 6, 6, 1, -1]
func ReplaceElements(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] < arr[j] {
				arr[i] = arr[j]
				arr[i+1] = arr[j+1]
				arr[i+2] = arr[j+2]
			}
		}
	}

	arr = arr[:len(arr)-1]
	arr = append(arr, -1)

	fmt.Println(arr)

	return arr
}
