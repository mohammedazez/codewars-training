package questions

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func NthEven(n int) int {
	if n == 1 {
		return 0
	}

	return n*2 - 2
}

var _ = Describe("Basic Tests", func() {
	It("Testing for 1", func() { Expect(NthEven(1)).To(Equal(0)) }) // 1 -> 0
	It("Testing for 2", func() { Expect(NthEven(2)).To(Equal(2)) }) // 2 -> 0, 2
	It("Testing for 3", func() { Expect(NthEven(3)).To(Equal(4)) }) // 3 -> 0, 2, 4
	It("Testing for 4", func() { Expect(NthEven(4)).To(Equal(4)) }) // 4 -> 0, 2, 4
	It("Testing for 5", func() { Expect(NthEven(5)).To(Equal(6)) }) // 5 -> 0, 2, 4, 6
	It("Testing for 100", func() { Expect(NthEven(100)).To(Equal(198)) })
	It("Testing for 1298734", func() { Expect(NthEven(1298734)).To(Equal(2597466)) })
})
