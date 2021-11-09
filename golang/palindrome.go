package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"strings"
)

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

var _ = Describe("Test Example", func() {
	It("tests basic strings", func() {
		Expect(IsPalindrome("a")).To(Equal(true))
		Expect(IsPalindrome("aba")).To(Equal(true))
		Expect(IsPalindrome("Abba")).To(Equal(true))
		Expect(IsPalindrome("hello")).To(Equal(false))
	})
})
