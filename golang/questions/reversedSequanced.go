package questions

func ReverseSeq(n int) []int {

	var results []int
	for i := n; i > 0; i-- {
		results = append(results, i)
	}

	return results
}
