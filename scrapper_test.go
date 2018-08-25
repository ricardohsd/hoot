package main

import (
	"testing"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/stretchr/testify/assert"
)

func TestScrapper_Fetch(t *testing.T) {
	s := &Scrapper{
		client: &testClient{},
	}

	tweetsC := s.Fetch("anyuser", 2)

	var tweets []twitter.Tweet

	for tweet := range tweetsC {
		tweets = append(tweets, tweet)
	}

	assert.Equal(t, 2, len(tweets))
	assert.Contains(t, tweets, twitter.Tweet{ID: 100, Text: "Some tweet"})
	assert.Contains(t, tweets, twitter.Tweet{ID: 200, Text: "Another tweet"})
}

type testClient struct{}

func (t *testClient) UserTimeline(user string, count int) ([]twitter.Tweet, error) {
	return []twitter.Tweet{
		twitter.Tweet{
			ID:   100,
			Text: "Some tweet",
		},
		twitter.Tweet{
			ID:   200,
			Text: "Another tweet",
		},
	}, nil
}
