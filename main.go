package main

import (
	"flag"
	"fmt"
	"golang-hn/models"
)

// It is possible to specify by flag the quantity of top stories you retrieve. By default it is 20.
// Also, using -cvs flag it is possible to specify the file path to save the result to a .csv file
func main() {
	csv := flag.Bool("csv", false, "Please specify a valid path for the file (/home/username/).")
	qtt := flag.Int("qtt", 20, "Please provide the quantity of topstories you want to get.")
	flag.Parse()

	if *csv {
		StoriesToCSV(*qtt)
	} else {
		PrintTopStories(*qtt)
	}
}

// StoriesToCSV function -
func StoriesToCSV(qtt int) {
	s := models.NewStory()
	path, e := s.CSVTopStories(qtt)
	if e != nil {
		fmt.Print(e.Error())
		return
	}
	fmt.Printf("A CSV file with %d stories was created on %s.", qtt, path)
}

// PrintTopStories function -
func PrintTopStories(qtt int) {
	s := models.NewStory()
	stories, e := s.GetTopStories(qtt)
	if e != nil {
		fmt.Print(e.Error())
		return
	}
	fmt.Print(stories)
}
