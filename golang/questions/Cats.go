package questions

import "fmt"

func Cats(start, finish int) int {
	a := finish - start
	fmt.Println(a)
	return a/3 + a%3
}
