package dynamicprogramming

import "fmt"

func MinPathSum(grid [][]int) int {
	var up, ln int
	for _, v := range grid[0] {
		up += v
	}
	for i := 1; i < len(grid); i++ {
		if len(grid) >= 3 {
			ln += grid[i][len(grid)-1]
		}
		if len(grid) <= 2 && len(grid[i]) >= 3 {
			outerLastElement := grid[len(grid)-1]
			innerLastElement := outerLastElement[len(outerLastElement)-1]
			ln = innerLastElement
		}
	}

	fmt.Println(up + ln)
	return up + ln
}
