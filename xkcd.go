package main

import (
	"fmt"

	"github.com/cevataykans/aws-lambda-go/feed"
)

var (
	rssFeed feed.RSS
	feedURL = "https://xkcd.com/rss.xml"
)

type EventRequest struct {
	Previous bool `json:"previous"`
}

type EventResponse struct {
	Title     string `json:"title"`
	URL       string `json:"url"`
	Published string `json:"published"`
}

func main() {
	fmt.Println("Hello world!")
}
