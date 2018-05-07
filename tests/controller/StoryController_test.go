package controller_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"golang-hn/configs"
	"golang-hn/controller"
	"reflect"
)

var _ = Describe("StoryController", func() {
	hnVars := configs.HackerNewsEnvV0()
	storyController := controller.StoryController(hnVars)

	BeforeEach(func() {
		hnVars = configs.HackerNewsEnvV0()
		storyController = controller.StoryController(hnVars)
	})

	Describe("Checking buffer length", func() {
		Context("Calling GetTopStories passing 20 as quantity", func() {
			It("should return an array of lenght 20", func() {
				stories, _ := storyController.GetTopStories(20)
				Expect(len(stories)).To(Equal(20))
			})
		})
	})

	Describe("Checking type of first element", func() {
		Context("Calling GetTopStoriesIds and testing the type of the element", func() {
			It("should be an integer", func() {
				stories, _ := storyController.GetTopStories(20)
				kidsType := reflect.TypeOf(stories[0].Kids).String()
				Expect(kidsType).To(Equal("[]int"))
			})
		})
	})
})
