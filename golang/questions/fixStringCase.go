package main

import (
	"fmt"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func solve(str string) string {
	// Your code here and happy coding!
	// fmt.Println(str)
	var temp string

	for i := 0; i < len(str); i++ {
		if string(str[i]) == strings.ToUpper(string(str[i])) {
			fmt.Println("big", string(str[i]))
		}
		if string(str[i]) == strings.ToLower(string(str[i])) {
			fmt.Println("small", string(str[i]))
		}
		// fmt.Println(string(str[i]))
	}
	return temp

}

// solve("coDe") = "code". Lowercase characters > uppercase. Change only the "D" to lowercase.
// solve("CODe") = "CODE". Uppercase characters > lowecase. Change only the "e" to uppercase.
// solve("coDE") = "code". Upper == lowercase. Change all to lowercase.

var _ = Describe("Sample Test Cases:", func() {
	It("Should return the correct values for the sample test cases!", func() {
		Expect(solve("a")).To(Equal("a"))
		Expect(solve("Z")).To(Equal("Z"))
		Expect(solve("coDe")).To(Equal("code"))
		Expect(solve("CODe")).To(Equal("CODE"))
		Expect(solve("coDE")).To(Equal("code"))
		Expect(solve("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")).To(Equal("abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz"))
	})
})
