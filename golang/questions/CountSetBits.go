package questions

// Given a set of an array of Integers,
// Find the smallest positive missing number possibility from an array of an integer as N, where N : 0 < N < 500000
// and then returns the numbers of bits set to 1 in the binary representation of the number N.
// For example given : [1,3,2,6,4,1,2] , the function should return 2 ,
// Find the smallest missing positive number from an array of integer (N) is 5
// The binary representation of N is 101, it contains 2 bits set to 1
// another example is given: [-1, -4], the function should return 1 (smallest positive integer), where N is 1 and binary representation of N is 1

// Solution is your solution code.
func Solution2(a []int) int {
	// Your code starts here.
	filter := make([]int, 0)
	for _, v := range a {
		if v > 0 {
			filter = append(filter, v)
		}
	}

	n := len(filter)
	for i := 0; i < n; i++ {
		if abs(filter[i])-1 < n && filter[abs(filter[i])-1] > 0 {
			filter[abs(filter[i])-1] = -filter[abs(filter[i])-1]
		}
	}
	for i := 0; i < n; i++ {
		if filter[i] > 0 {
			missing := i + 1
			bitCount := countSetBits(missing)
			return bitCount
		}
	}
	return countSetBits(n + 1)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func countSetBits(n int) int {
	count := 0
	for n > 0 {
		count += n & 1
		n >>= 1
	}
	return count
}
