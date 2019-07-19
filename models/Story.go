package models

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
	"time"
)

type story struct {
	By          string `bson:"by",json:"by"`
	Descendants int    `bson:"descendants",json:"descendants"`
	Id          int    `bson:"id",json:"id"`
	Kids        []int  `bson:"kids",json:"kids"`
	Score       int    `bson:"score",json:"score"`
	Time        int    `bson:"time",json:"time"`
	Title       string `bson:"title",json:"title"`
	Type        string `bson:"type",json:"type"`
	Url         string `bson:"url",json:"url"`
}

// NewStory function -
func NewStory() *story {
	return &story{}
}

// Get the specified quantity of topstories ids
func (this *story) GetTopStoriesIds(qtt int) (storiesIds []int, e error) {
	var url = "https://hacker-news.firebaseio.com/v0/topstories.json?print=pretty&orderBy=\"$key\"&limitToFirst=%d"

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		var res *http.Response

		res, e = http.Get(fmt.Sprintf(url, qtt)) // *http.Response
		if e != nil {
			return
		}

		var contents []uint8
		contents, e = ioutil.ReadAll(res.Body)
		if e != nil {
			return
		}

		e = json.Unmarshal(contents, &storiesIds)
		defer res.Body.Close()
		defer wg.Done()
	}()
	wg.Wait()
	return
}

// Get a topstory from the specified id
func (this *story) GetTopStories(qtt int) (stories []story, e error) {
	var url = "https://hacker-news.firebaseio.com/v0/item/%d.json?print=pretty"

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		var ids []int

		ids, e = this.GetTopStoriesIds(qtt)
		if e != nil {
			return
		}

		var res *http.Response
		for _, id := range ids {
			res, e = http.Get(fmt.Sprintf(url, id))
			if e != nil {
				return
			}

			var item []uint8
			item, e = ioutil.ReadAll(res.Body)
			if e != nil {
				return
			}

			e = json.Unmarshal(item, this)
			if e != nil {
				return
			}

			stories = append(stories, *this)
		}
		defer res.Body.Close()
		defer wg.Done()
	}()
	wg.Wait()
	return
}

//
func (this *story) CSVTopStories(qtt int) (path string, e error) {
	// Create a csv file
	path = fmt.Sprintf("csv/%dtopstories-%s.csv", qtt, time.Now().Format("20060102150405"))
	osFile, err := os.Create(path)
	if err != nil {
		return path, err
	}
	defer osFile.Close()

	stories, err := this.GetTopStories(qtt)
	if err != nil {
		return path, err
	}

	w := csv.NewWriter(osFile)
	for _, story := range stories {
		var record []string
		record = append(record, story.By)
		record = append(record, fmt.Sprint(story.Descendants))
		record = append(record, fmt.Sprint(story.Id))
		record = append(record, fmt.Sprint(story.Kids))
		record = append(record, fmt.Sprint(story.Score))
		record = append(record, fmt.Sprint(story.Time))
		record = append(record, story.Title)
		record = append(record, story.Type)
		record = append(record, story.Url)
		w.Write(record)
	}
	w.Flush()

	return
}
