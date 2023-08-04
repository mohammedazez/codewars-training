package questions

import "strings"

type MyString string

func (s MyString) IsUpperCase() bool {
	// Your code here!
	return strings.ToUpper(string(s)) == string(s)
}
