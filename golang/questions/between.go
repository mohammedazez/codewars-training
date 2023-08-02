package questions

func Between(a, b int) []int {
	// your code here
	var results []int
	for i := a; i <= b; i++ {
		results = append(results, i)
	}

	return results
}
