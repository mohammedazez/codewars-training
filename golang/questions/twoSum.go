package questions

// 2, 7, 11, 15 					-> FIRST LOOPING
// 7 ,11, 15     11, 15      15 	-> SECOND LOOPING
//

func TwoSum(nums []int, target int) []int {
	var res []int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == target-nums[i] {
				res = append(res, i, j)
			}
		}
	}
	return res
}
