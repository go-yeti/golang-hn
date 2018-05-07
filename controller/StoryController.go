package controller

import (
	"encoding/csv"
	"fmt"
	"golang-hn/configs"
	"golang-hn/dao"
	"golang-hn/model"
	"golang-hn/utils"
	"os"
	//	"reflect"
	"time"
)

type storyController struct {
	envVar configs.EnvInterface
}

func StoryController(vars configs.EnvInterface) *storyController {
	return &storyController{envVar: vars}
}

// Get the specified quantity of topstories items
func (sc *storyController) GetTopStories(qtt int) ([]model.Story, error) {
	storyDAO := dao.NewStoryDAO(sc.envVar)

	var stories []model.Story
	var story model.Story
	var err error

	topStoriesIds, err := storyDAO.GetTopStoriesIds(qtt)
	if err != nil {
		return stories, err
	}

	for _, id := range topStoriesIds {
		story, err = storyDAO.GetTopStory(id)
		if err != nil {
			return stories, err
		}
		stories = append(stories, story)
	}

	return stories, err
}

// save topstories to a csv file. The app is waiting for the path only, the name
// of the file, for security reasons, is automatically generated, based on timestamp
// and the number of stories, avoiding conflicts and file overwriting
func (sc *storyController) TopStories2CSV(qtt int, path string) (string, error) {
	currTime := time.Now()

	// Create a csv file
	filePath, err := os.Create(utils.StringConcat(path, fmt.Sprintf("%d", qtt), "_topstories-", currTime.Format("20060102150405"), ".csv"))
	if err != nil {
		return "", err
	}
	defer filePath.Close()

	stories, err := sc.GetTopStories(qtt)
	if err != nil {
		return "", err
	}

	w := csv.NewWriter(filePath)
	for _, story := range stories {
		var record []string
		record = append(record, story.By)
		record = append(record, fmt.Sprintf("%d", story.Descendants))
		record = append(record, fmt.Sprintf("%d", story.Id))
		record = append(record, utils.IntArray2String(story.Kids, " "))
		record = append(record, fmt.Sprintf("%d", story.Score))
		record = append(record, fmt.Sprintf("%d", story.Time))
		record = append(record, story.Title)
		record = append(record, story.Type)
		record = append(record, story.Url)
		w.Write(record)
	}
	w.Flush()

	return "The file was successfully saved", err
}
