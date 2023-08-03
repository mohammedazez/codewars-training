package questions

func Hero(bullets, dragons int) bool {
	// your code
	var results bool
	sum := dragons * 2
	if bullets == sum || bullets >= sum {
		results = true
	}

	return results
}

// simple case
func Hero2(bullets, dragons int) bool {
	return bullets >= 2*dragons
}
