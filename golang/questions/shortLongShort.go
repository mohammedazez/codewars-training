package questions

func Solution(a, b string) string {
	// TODO: Your code here
	var results string
	if len(a) > len(b) {
		results = b + a + b
	} else {
		results = a + b + a
	}

	return results
}
