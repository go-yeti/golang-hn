package dao

import (
	"encoding/json"
	"fmt"
	"golang-hn/configs"
	"golang-hn/model"
	"golang-hn/utils"
	"io/ioutil"
	"net/http"
	"runtime"
	"sync"
	"time"
)

var httpClient = &http.Client{Timeout: 10 * time.Second}

type storiesChnRtn struct {
	Result interface{}
	Err    error
	Info   interface{}
}

type storyDAO struct {
	envVar configs.EnvInterface
}

func NewStoryDAO(vars configs.EnvInterface) *storyDAO {
	return &storyDAO{envVar: vars}
}

// Get the specified quantity of topstories ids
func (sd *storyDAO) GetTopStoriesIds(qtt int) ([]int, error) {
	var wg sync.WaitGroup
	vars := sd.envVar.GetVars()
	url := utils.StringConcat(vars["baseUrl"], vars["topStories"])

	var storiesIds []int

	storiesChn := make(chan storiesChnRtn)

	wg.Add(1)
	go func() {
		r, err := httpClient.Get(url)
		defer r.Body.Close()
		if err != nil {
			storiesChn <- storiesChnRtn{
				Result: storiesIds,
				Err:    err,
			}
			runtime.Goexit()
		}

		response, err := ioutil.ReadAll(r.Body)
		if err != nil {
			storiesChn <- storiesChnRtn{
				Result: storiesIds,
				Err:    err,
			}
			runtime.Goexit()
		}

		err = json.Unmarshal(response, &storiesIds)
		storiesChn <- storiesChnRtn{
			Result: storiesIds[:qtt],
			Err:    err,
		}
		defer wg.Done()
	}()
	resp := <-storiesChn
	defer close(storiesChn)
	wg.Wait()
	return resp.Result.([]int), resp.Err
}

// Get a topstory from the specified id
func (sd *storyDAO) GetTopStory(id int) (model.Story, error) {
	var wg sync.WaitGroup
	vars := sd.envVar.GetVars()
	url := utils.StringConcat(vars["baseUrl"], fmt.Sprintf(vars["item"], id))

	var story model.Story

	storiesChn := make(chan storiesChnRtn)

	go func() {
		r, err := httpClient.Get(url)
		if err != nil {
			storiesChn <- storiesChnRtn{
				Result: story,
				Err:    err,
			}
			runtime.Goexit()
		}

		defer r.Body.Close()
		response, err := ioutil.ReadAll(r.Body)
		if err != nil {
			storiesChn <- storiesChnRtn{
				Result: story,
				Err:    err,
			}
			runtime.Goexit()
		}

		err = json.Unmarshal(response, &story)
		if err != nil {
			storiesChn <- storiesChnRtn{
				Result: story,
				Err:    err,
			}
		}
		defer wg.Done()
	}()
	resp := <-storiesChn
	defer close(storiesChn)
	wg.Wait()
	return resp.Result.(model.Story), resp.Err
}

// Get a topstory from the specified id
//func (sd *storyDAO) GetTopStory(id int) (model.Story, error) {
//	vars := sd.envVar.GetVars()
//	url := utils.StringConcat(vars["baseUrl"], fmt.Sprintf(vars["item"], id))
//	var story model.Story
//
//	r, err := httpClient.Get(url)
//	if err != nil {
//		return story, err
//	}
//
//	defer r.Body.Close()
//	response, err := ioutil.ReadAll(r.Body)
//	if err != nil {
//		return story, err
//	}
//
//	err = json.Unmarshal(response, &story)
//	if err != nil {
//		return story, err
//	}
//
//	return story, err
//}

// Get a newstory from the specified id
func (sd *storyDAO) GetNewStory(id int) {}

// Get a beststory from the specified id
func (sd *storyDAO) GetBestStory(id int) {}
