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

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	if len(dbUser) == 0 || len(dbPassword) == 0 || len(dbName) == 0 {
		fmt.Println("env variables DB_USER, DB_PASSWORD, DB_NAME must be set")
		os.Exit(1)
	}

	client := NewTwitterClient(consumerKey, consumerSecret, token, tokenSecret)

	storage := NewStorage(dbUser, dbPassword, dbName)
	scrapper := NewScrapper(client)

	tweets := scrapper.Fetch(user, amount)
	storage.Ingest(tweets)

	fmt.Println("closing...")
}
