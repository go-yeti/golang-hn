package dao_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"golang-hn/configs"
	"golang-hn/dao"
	"reflect"
)

var _ = Describe("StoryDAO", func() {

	hnVars := configs.HackerNewsEnvV0()
	storyDAO := dao.NewStoryDAO(hnVars)

	BeforeEach(func() {
		hnVars = configs.HackerNewsEnvV0()
		storyDAO = dao.NewStoryDAO(hnVars)
	})

	Describe("Checking buffer length", func() {
		Context("Calling GetTopStoriesIds passing 20 as quantity", func() {
			It("should return an array of lenght 20", func() {
				storiesIds, _ := storyDAO.GetTopStoriesIds(20)
				Expect(len(storiesIds)).To(Equal(20))
			})
		})
	})

	Describe("Checking type of first element", func() {
		Context("Calling GetTopStoriesIds and testing the type of the element", func() {
			It("should be an integer", func() {
				storyId, _ := storyDAO.GetTopStoriesIds(1)
				storyIdType := reflect.TypeOf(storyId[0]).String()
				Expect(storyIdType).To(Equal("int"))
			})
		})
	})

	Describe("Checking story item", func() {
		Context("Calling GetTopStory to check the returned element", func() {
			storyId, _ := storyDAO.GetTopStoriesIds(1)
			storyItem, _ := storyDAO.GetTopStory(storyId[0])
			It("should be a string", func() {
				byType := reflect.TypeOf(storyItem.By).String()
				Expect(byType).To(Equal("string"))
			})
			It("should be a string", func() {
				timeType := reflect.TypeOf(storyItem.Time).String()
				Expect(timeType).To(Equal("int"))
			})
			It("should be a []int", func() {
				kidsType := reflect.TypeOf(storyItem.Kids).String()
				Expect(kidsType).To(Equal("[]int"))
			})
		})
	})
})
