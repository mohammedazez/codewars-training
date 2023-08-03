package questions

func Xor(a, b bool) bool {
	// your code here:
	var one int
	var two int
	var results bool

	if a {
		one = 1
	} else {
		one = 0
	}

	if b {
		two = 1
	} else {
		two = 0
	}

	xor := one ^ two
	if xor == 1 {
		results = true
	} else {
		results = false
	}

	return results
}

// simple case
func Xor2(a, b bool) bool {
	return a != b
}
