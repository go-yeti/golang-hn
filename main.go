package main

import (
	"flag"
	"fmt"
	"golang-hn/configs"
	"golang-hn/controller"
)

// It is possible to specify by flag the quantity of top stories you retrieve. By default it is 20.
// Also, using -cvs flag it is possible to specify the file path to save the result to a .csv file
func main() {
	var csvFilePath string
	flag.StringVar(&csvFilePath, "csv", "", "Please specify a valid path for the file (/home/username/).")
	var qttStories int
	flag.IntVar(&qttStories, "tsqtt", 20, "Please provide the quantity of topstories you want to get.")
	flag.Parse()

	// get the environment variables
	hnEnv := configs.HackerNewsEnvV0()

	if len(csvFilePath) > 0 {
		SaveTopStories2CSV(hnEnv, qttStories, csvFilePath)
	} else {
		GetTopStories(hnEnv, qttStories)
	}
}

// Using dependency injection to avoid rewriting new code in case of API version change
func GetTopStories(hnEnv configs.EnvInterface, qtt int) {
	stories, err := controller.StoryController(hnEnv).GetTopStories(qtt)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(stories)
	}
}

// Save topstories to a specified .csv file
func SaveTopStories2CSV(hnEnv configs.EnvInterface, qtt int, filePath string) {
	msg, err := controller.StoryController(hnEnv).TopStories2CSV(qtt, filePath)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(msg)
	}
}
