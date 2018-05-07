package utils_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"golang-hn/utils"
)

var _ = Describe("StringConcat", func() {

	Describe("Checking string concatenation function", func() {
		Context("Calling StringConcat passing 'concatenated', ' string'", func() {
			It("should return 'concatenated string", func() {
				ctnStr := utils.StringConcat("concatenated", " string")
				Expect(ctnStr).To(Equal("concatenated string"))
			})
		})
	})

	Describe("Checking the int array to string conversion", func() {
		Context("Calling IntArray2String", func() {
			It("should be '2,4,6,8'", func() {
				evenNs := []int{2, 4, 6, 8}
				Expect(utils.IntArray2String(evenNs, ",")).To(Equal("2,4,6,8"))
			})
		})
	})
})
