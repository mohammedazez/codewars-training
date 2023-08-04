package questions

func GetCount(str string) (count int) {
	// Enter solution fihere
	key := "aiueo"
	var result []string
	for _, v := range key {
		for _, s := range str {
			if string(v) == string(s) {
				result = append(result, string(s))
			}
		}
	}
	return len(result)
}

// simple case
func GetCount2(str string) (count int) {
	for _, c := range str {
		switch c {
		case 'a', 'e', 'i', 'o', 'u':
			count++
		}
	}
	return count
}
