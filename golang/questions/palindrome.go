package questions

import (
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func IsPalindrome(str string) bool {

	lowerCase := strings.ToLower(str)
	for i := 0; i < len(string(lowerCase))/2; i++ {
		// fmt.Println(string(lowerCase[i]))
		// fmt.Println(string(lowerCase[len(lowerCase)-1-i]))
		if string(lowerCase[i]) != string(lowerCase[len(lowerCase)-1-i]) {
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
