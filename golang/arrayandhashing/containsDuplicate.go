package arrayandhashing

func ContainsDuplicate(nums []int) bool {
	Hashset := make(map[int]bool)
	for _, v := range nums {
		if Hashset[v] {
			return true
		}
		Hashset[v] = true
	}
	return false
}
