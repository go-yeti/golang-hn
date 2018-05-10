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

	storiesChn := make(chan model.ChnResult)
	defer close(storiesChn)

	wg.Add(1)
	go func() {
		response, err := httpClient.Get(url)
		defer response.Body.Close()
		defer wg.Done()

		if err != nil {
			storiesChn <- model.ChnResult{
				Result: storiesIds,
				Err:    err,
			}
			runtime.Goexit()
		}

		respIO, err := ioutil.ReadAll(r.Body)
		if err != nil {
			storiesChn <- model.ChnResult{
				Result: storiesIds,
				Err:    err,
			}
			runtime.Goexit()
		}

		err = json.Unmarshal(respIO, &storiesIds)
		storiesChn <- model.ChnResult{
			Result: storiesIds[:qtt],
			Err:    err,
		}
	}()
	resChn := <-storiesChn
	wg.Wait()
	return resChn.Result.([]int), resChn.Err
}

// Get a topstory from the specified id
func (sd *storyDAO) GetTopStory(id int) (model.Story, error) {
	var wg sync.WaitGroup
	vars := sd.envVar.GetVars()
	url := utils.StringConcat(vars["baseUrl"], fmt.Sprintf(vars["item"], id))

	var story model.Story

	storiesChn := make(chan model.ChnResult)
	defer close(storiesChn)

	wg.Add(1)
	go func() {
		response, err := httpClient.Get(url)
		defer response.Body.Close()
		defer wg.Done()

		if err != nil {
			storiesChn <- model.ChnResult{
				Result: story,
				Err:    err,
			}
			runtime.Goexit()
		}

		respIO, err := ioutil.ReadAll(r.Body)
		if err != nil {
			storiesChn <- model.ChnResult{
				Result: story,
				Err:    err,
			}
			runtime.Goexit()
		}

		err = json.Unmarshal(respIO, &story)
		storiesChn <- model.ChnResult{
			Result: story,
			Err:    err,
		}
	}()
	resChn := <-storiesChn
	wg.Wait()
	return resChn.Result.(model.Story), resChn.Err
}

// Get a newstory from the specified id
func (sd *storyDAO) GetNewStory(id int) {}

// Get a beststory from the specified id
func (sd *storyDAO) GetBestStory(id int) {}
