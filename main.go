package main

import (
	"flag"
	"fmt"
	"golang-hn/configs"
	"golang-hn/controller"
)

func main() {
	var csvFilePath string
	flag.StringVar(&csvFilePath, "csv", "", "Please specify a valid path for the file (/home/username/).")
	var qttStories int
	flag.IntVar(&qttStories, "qtt", 20, "Please specify a valid path for the file (/home/username/).")
	flag.Parse()

	// get the environment variables
	hnEnv := configs.HackerNewsEnvV0()

	if len(csvFilePath) > 0 {
		SaveTopStories2CSV(hnEnv, qttStories, csvFilePath)
	} else {
		Get20TopStories(hnEnv, qttStories)
	}
}

// Using dependency injection to avoid rewriting new code in case of API version change
func Get20TopStories(hnEnv configs.EnvInterface, qtt int) {
	stories, err := controller.StoryController(hnEnv).GetTopStories(qtt)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(stories)
	}
}

func SaveTopStories2CSV(hnEnv configs.EnvInterface, qtt int, filePath string) {
	msg, err := controller.StoryController(hnEnv).TopStories2CSV(qtt, filePath)
	if err != nil {
		panic(err)
	} else {
		fmt.Println(msg)
	}
}
