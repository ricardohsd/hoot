package main

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

type Tweeter interface {
	UserTimeline(user string, count int) ([]twitter.Tweet, error)
}

type twitterClient struct {
	client *twitter.Client
}

func NewTwitterClient(cKey, cSecret, t, tSecret string) *twitterClient {
	config := oauth1.NewConfig(cKey, cSecret)
	token := oauth1.NewToken(t, tSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	return &twitterClient{
		client: twitter.NewClient(httpClient),
	}
}

func (t *twitterClient) UserTimeline(user string, count int) ([]twitter.Tweet, error) {
	userTimelineParams := &twitter.UserTimelineParams{ScreenName: user, Count: count}
	tweets, _, err := t.client.Timelines.UserTimeline(userTimelineParams)
	return tweets, err
}
