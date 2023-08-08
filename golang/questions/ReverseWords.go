package questions

import (
	"strings"
)

func ReverseWords(str string) string {
	fld := strings.Fields(str)
	var rev []string
	for _, n := range fld {
		rev = append([]string{n}, rev...)
	}
	return strings.Join(rev, " ") // reverse those words
}
