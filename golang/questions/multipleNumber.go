package questions

func FindMultiples(integer, limit int) []int {
	// Your code here!
	var result []int

	for i := 1; i <= limit; i++ {
		sum := integer + (i-1)*integer
		if sum <= limit {
			result = append(result, sum)
		}
	}

	return result
}
