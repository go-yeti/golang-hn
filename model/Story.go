package model

type Story struct {
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
