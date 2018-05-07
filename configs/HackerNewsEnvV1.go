package configs

type hackerNewsEnvV1 struct {
	Vars map[string]string
}

func HackerNewsEnvV1() *hackerNewsEnvV1 {
	return &hackerNewsEnvV1{
		Vars: map[string]string{
			"baseUrl":     "https://hacker-news.firebaseio.com/v0/",
			"topStories":  "topstories.json",
			"newStories":  "newstories.json",
			"bestStories": "beststories.json",
			"item":        "item/%d.json",
		},
	}
}

func (hn *hackerNewsEnvV1) GetVars() map[string]string {
	return hn.Vars

}

func (hn *hackerNewsEnvV1) SetVars(newVars map[string]string) {
	hn.Vars["baseUrl"] = newVars["baseUrl"]
}
