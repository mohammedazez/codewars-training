package questions

// []int{2, 7, 11, 15}, 9
// 2, 7, 11, 15 					-> FIRST LOOPING
// 7 ,11, 15     11, 15      15 	-> SECOND LOOPING
// apakah 7 sama dengan 9 - 2 --> 7 == 7 jadi iya
// apakah 11 sama dengan 9 - 7 --> 2 == 11 jadi tidak
// apakah 15 sama dengan 9 - 15 --> 6 == 15 jadi tidak

func TwoSum(nums []int, target int) []int {
	var res []int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			// if nums[j] == target-nums[i] {
			// 	res = append(res, i, j)
			// }

			if nums[i]+nums[j] == target {
				res = append(res, i, j)
			}
		}
	}
	return res
}
