package main

import (
	"fmt"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/go-pg/pg"
)

type Tweet struct {
	ID        string
	Username  string
	TweetID   string
	Content   string
	PostedAt  string
	CreatedAt time.Time
}

type storage struct {
	db *pg.DB
}

func NewStorage(user, password, dbName string) *storage {
	db := pg.Connect(&pg.Options{
		User:     user,
		Password: password,
		Database: dbName,
	})

	return &storage{
		db: db,
	}
}

func (s *storage) Ingest(tweetsC <-chan twitter.Tweet) {
	for tweet := range tweetsC {
		err := s.insert(tweet)
		if err != nil {
			fmt.Printf("Error to save tweet %v, %v\n", tweet.ID, err)
		}

		count, err := s.db.Model(&Tweet{}).Count()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%v records saved\n", count)
	}
}

func (s *storage) insert(tweet twitter.Tweet) error {
	user := tweet.User

	var userName string
	if user != nil {
		userName = user.ScreenName
	}

	return s.db.Insert(&Tweet{
		Username: userName,
		TweetID:  fmt.Sprintf("%v", tweet.ID),
		Content:  tweet.Text,
		PostedAt: tweet.CreatedAt,
	})
}
