package hoot

import (
	"fmt"
	"os"
	"testing"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/go-pg/pg"
	"github.com/stretchr/testify/assert"
)

func TestStorage_Ingest(t *testing.T) {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := fmt.Sprintf("%v_test", os.Getenv("DB_NAME"))

	db := pg.Connect(&pg.Options{
		User:     dbUser,
		Password: dbPassword,
		Database: dbName,
	})

	s := &storage{
		db: db,
	}

	tweetsC := make(chan twitter.Tweet)

	go func() {
		tweetsC <- twitter.Tweet{
			ID:   100,
			Text: "Some tweet",
			User: &twitter.User{
				ScreenName: "userName",
			},
			CreatedAt: "Thu Aug 23 20:27:42 +0000 2018",
		}
		tweetsC <- twitter.Tweet{
			ID:   200,
			Text: "Another tweet",
			User: &twitter.User{
				ScreenName: "userName",
			},
			CreatedAt: "Thu Aug 23 20:27:42 +0000 2018",
		}
		close(tweetsC)
	}()

	s.Ingest(tweetsC)

	var tweets []Tweet
	err := db.Model(&tweets).Select()
	assert.Nil(t, err, "Failed to fetch tweets from database")

	assert.Equal(t, 2, len(tweets))
	assertTweet(t, Tweet{
		TweetID:  "100",
		Content:  "Some tweet",
		Username: "userName",
		PostedAt: "Thu Aug 23 20:27:42 +0000 2018",
	}, tweets[0])
	assertTweet(t, Tweet{
		TweetID:  "200",
		Content:  "Another tweet",
		Username: "userName",
		PostedAt: "Thu Aug 23 20:27:42 +0000 2018",
	}, tweets[1])
}

func assertTweet(t *testing.T, expected Tweet, tweet Tweet) {
	assert.Equal(t, expected.TweetID, tweet.TweetID)
	assert.Equal(t, expected.Content, tweet.Content)
	assert.Equal(t, expected.Username, tweet.Username)
	assert.Equal(t, expected.PostedAt, tweet.PostedAt)
}
