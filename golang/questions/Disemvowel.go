package questions

import (
	"strings"
)

func Disemvowel(comment string) (res string) {

	for i := 0; i < len(comment); i++ {
		if !strings.ContainsAny(strings.ToLower(string(comment[i])), "aiueo") {
			res += string(comment[i])
		}
	}
	return res
}
