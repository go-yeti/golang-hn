package dao

import (
	"encoding/json"
	"fmt"
	"golang-hn/configs"
	"golang-hn/model"
	"golang-hn/utils"
	"io/ioutil"
	"net/http"
	"time"
)

var httpClient = &http.Client{Timeout: 10 * time.Second}

type storyDAO struct {
	envVar configs.EnvInterface
}

func NewStoryDAO(vars configs.EnvInterface) *storyDAO {
	return &storyDAO{envVar: vars}
}

// Get the specified quantity of topstories ids
func (sd *storyDAO) GetTopStoriesIds(qtt int) ([]int, error) {
	vars := sd.envVar.GetVars()
	var storiesIds []int

	url := utils.StringConcat(vars["baseUrl"], vars["topStories"])

	r, err := httpClient.Get(url)
	if err != nil {
		return storiesIds, err
	}
	defer r.Body.Close()

	response, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return storiesIds, err
	}

	err = json.Unmarshal(response, &storiesIds)

	return storiesIds[:qtt], err
}

// Get a topstory from the specified id
func (sd *storyDAO) GetTopStory(id int) (model.Story, error) {
	vars := sd.envVar.GetVars()
	url := utils.StringConcat(vars["baseUrl"], fmt.Sprintf(vars["item"], id))
	var story model.Story

	r, err := httpClient.Get(url)
	if err != nil {
		return story, err
	}

	defer r.Body.Close()
	response, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return story, err
	}

	err = json.Unmarshal(response, &story)
	if err != nil {
		return story, err
	}

	return story, err
}

// Get a newstory from the specified id
func (sd *storyDAO) GetNewStory(id int) {}

// Get a beststory from the specified id
func (sd *storyDAO) GetBestStory(id int) {}
