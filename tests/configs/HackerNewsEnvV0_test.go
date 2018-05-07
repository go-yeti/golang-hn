package configs_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"golang-hn/configs"
	"reflect"
)

var _ = Describe("HackerNewsEnvV0", func() {
	hnVars := configs.HackerNewsEnvV0()
	vars := hnVars.GetVars()

	BeforeEach(func() {
		hnVars = configs.HackerNewsEnvV0()
		vars = hnVars.GetVars()
	})

	Describe("Checking the GetVars() method", func() {
		Context("Checking the type of baseUrl element", func() {
			It("Should be a string", func() {
				Expect(reflect.TypeOf(vars["baseUrl"]).String()).To(Equal("string"))
			})
		})

		Context("Checking the value of baseUrl element", func() {
			It("should be 'https://hacker-news.firebaseio.com/v0/'", func() {
				Expect(vars["baseUrl"]).To(Equal("https://hacker-news.firebaseio.com/v0/"))
			})
		})
	})

	Describe("Checking the SetVars() method", func() {
		Context("Checking the value of baseUrl element", func() {
			It("Should be a 'https://github.com/jamesmallon'", func() {
				hnVars.SetVars(map[string]string{
					"baseUrl":     "https://github.com/jamesmallon",
					"topStories":  "topStories",
					"newStories":  "newStories",
					"bestStories": "bestStories",
					"item":        "item",
				})
				vars = hnVars.GetVars()
				Expect(vars["baseUrl"]).To(Equal("https://github.com/jamesmallon"))
			})
		})
	})
})
