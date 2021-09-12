package main

import "strings"

func IsPalindrome(str string) bool {
	// fmt.Println("str", str)
	lowerCase := strings.ToLower(str)
	for i := 0; i < len(lowerCase)/2; i++ {
		// fmt.Println("first", string(str[i]))
		// fmt.Println("last", string(str[len(str)-i-1]))
		if string(lowerCase[i]) != string(lowerCase[len(lowerCase)-i-1]) {
			return false
		}
	}
	return true
}
