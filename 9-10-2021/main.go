package main

import (
	"fmt"
	"strings"
)

func MakeUpperCase(word string) string {
	result := strings.ToUpper(word)
	return result
}

func main() {
	fmt.Println(MakeUpperCase("hello"))
}

// func MakeUpperCase(str string) string {
// 	return ""
// }
