package questions

func Grow(arr []int) (res int) {
	res = 1
	for i := 0; i < len(arr); i++ {
		res *= arr[i]
	}
	return
}
