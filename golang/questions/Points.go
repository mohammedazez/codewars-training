package questions

func Points(games []string) int {
	// your code here!
	var total int
	for i := 0; i < len(games); i++ {
		if string(games[i][0]) > string(games[i][2]) {
			total += 3
		}

		if string(games[i][0]) == string(games[i][2]) {
			total += 1
		}
	}

	return total
}
