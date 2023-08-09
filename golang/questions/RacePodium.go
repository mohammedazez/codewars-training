package questions

func RacePodium(blocks int) [3]int {
	if blocks == 7 {
		return [3]int{2, 4, 1}
	}
	x := (blocks + 5) / 3
	return [3]int{x - 1, x, blocks - 2*x + 1}
}
