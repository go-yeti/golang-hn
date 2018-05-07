package configs

type hackerNewsEnvV0 struct {
	Vars map[string]string
}

func HackerNewsEnvV0() *hackerNewsEnvV0 {
	return &hackerNewsEnvV0{
		Vars: map[string]string{
			"baseUrl":     "https://hacker-news.firebaseio.com/v0/",
			"topStories":  "topstories.json",
			"newStories":  "newstories.json",
			"bestStories": "beststories.json",
			"item":        "item/%d.json",
		},
	}
}

// Retrieve the environment variables
func (hn *hackerNewsEnvV0) GetVars() map[string]string {
	return hn.Vars

}

// Set the environment variables
func (hn *hackerNewsEnvV0) SetVars(newVars map[string]string) {
	hn.Vars["baseUrl"] = newVars["baseUrl"]
	hn.Vars["topStories"] = newVars["topStories"]
	hn.Vars["newStories"] = newVars["newStories"]
	hn.Vars["bestStories"] = newVars["bestStories"]
	hn.Vars["item"] = newVars["item"]
}
