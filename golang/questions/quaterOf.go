package questions

func QuarterOf(month int) int {
	// your code here
	var result int

	switch month {
	case 1, 2, 3:
		result = 1
		return result
	case 4, 5, 6:
		result = 2
		return result
	case 7, 8, 9:
		result = 3
		return result
	case 10, 11, 12:
		result = 4
		return result
	}

	return result
}
