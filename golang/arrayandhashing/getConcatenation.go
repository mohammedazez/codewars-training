package arrayandhashing

func GetConcatenation(nums []int) []int {
	s := nums

	for i := 0; i < len(nums); i++ {
		s = append(s, nums[i])
	}

	return s
}
