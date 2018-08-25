package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	consumerKey := os.Getenv("CONSUMER_KEY")
	consumerSecret := os.Getenv("CONSUMER_SECRET")
	token := os.Getenv("TOKEN")
	tokenSecret := os.Getenv("TOKEN_SECRET")

	var amount int
	flag.IntVar(&amount, "amount", 5, "amount of tweets to fetch")

	var user string
	flag.StringVar(&user, "user", "", "twitter username")

	flag.Parse()

	if len(user) == 0 {
		fmt.Println("-user= must be present")
		os.Exit(1)
	}

	client := NewTwitterClient(consumerKey, consumerSecret, token, tokenSecret)

	scrapper := NewScrapper(client)

	tweets := scrapper.Fetch(user, amount)
	for tweet := range tweets {
		fmt.Printf("Tweet ID: %v, Text: %v\n", tweet.ID, tweet.Text)
	}

	fmt.Println("closing...")
}
