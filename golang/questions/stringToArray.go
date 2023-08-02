package questions

import (
	"strings"
)

func StringToArray(str string) []string {
	// your code here
	fields := strings.Fields(str)

	return fields
}
