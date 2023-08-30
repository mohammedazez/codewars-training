package questions

import (
	"strings"
	"unicode/utf8"
)

// bar bar case
func ToJadenCase(str string) string {
	fields := strings.Fields(str)
	var res []string
	for i := 0; i < len(fields); i++ {
		_, g := utf8.DecodeRuneInString(fields[i])
		rel := strings.ToUpper(string(fields[i][0])) + fields[i][g:]
		res = append(res, rel)
	}

	return strings.Join(res, " ")
}

// "most trees are blue" -> "Most Trees Are Blue"
// ["most", "tress", "are", "blue"]
// "Most" "Tress" "Are" "Blue"
// "ost" "ress" "re" "lue" -> delete first index
// M T A B -> Uppercase first index
// "Most" "Tress" "Are" "Blue"
// "Most Trees Are Blue"

// simple case
func ToJadenCase2(str string) string {
	return strings.Title(str)
}
