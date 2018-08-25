package main

import (
	"github.com/dghubble/go-twitter/twitter"
)

type Scrapper struct {
	client Tweeter
}

func NewScrapper(client Tweeter) *Scrapper {
	return &Scrapper{
		client: client,
	}
}

func (s *Scrapper) Fetch(user string, count int) <-chan twitter.Tweet {
	tweetsC := make(chan twitter.Tweet)

	go func() {
		tweets, err := s.client.UserTimeline(user, count)
		if err != nil {
			panic(err)
		}

		for _, tweet := range tweets {
			tweetsC <- tweet
		}

		close(tweetsC)
	}()

	return tweetsC
}
