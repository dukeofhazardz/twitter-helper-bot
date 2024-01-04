package main

import (
	"fmt"
	"os"
	"time"

	helper "twitter-helper-bot/twitterhelper"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init() {

	log.Out = os.Stdout

	// You could set this to any `io.Writer` such as a file
	file, err := os.OpenFile("logs.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.Out = file
	} else {
		log.Info("Failed to log to file, using default stderr")
	}

}

func main() {
	// Calculate start time for bot server
	startTime := time.Now().Local()
	starttimeStampString := startTime.Format("2006-01-02 15:04:05")
	unixNano := startTime.UnixNano()
	umillisec := unixNano / 1000000
	fmt.Println("+++++++++++++++++++++++++++++++++++++++")
	fmt.Println("Starting the Twitter-Helper-Bot application.....", starttimeStampString)
	fmt.Println("+++++++++++++++++++++++++++++++++++++++")

	creds := helper.TwitterCredentials{
		Consumerkey:       os.Getenv("TWITTER_CONSUMER_KEY"),
		ConsumerSecret:    os.Getenv("TWITTER_CONSUMER_SECRET"),
		AccessToken:       os.Getenv("TWITTER_ACCESS_TOKEN"),
		AccessTokenSecret: os.Getenv("TWITTER_ACCESS_SECRET"),
	}

	client, err := helper.Twitter(&creds)

	if err != nil {
		log.Error("Error in Twitter client", client)
		//log.Println(err)
	}
	/*	*/
	TrendListName := helper.TopTrending(client)
	SearchQueries := helper.Searchqueries(TrendListName, client)
	TweetStatusIds := helper.FilterTweetstatusId(SearchQueries, client)
	TweetStatusDescription := helper.RetweetingTweet(TweetStatusIds, client)

	helper.PostTweet("Let's Connect Together", client)

	fmt.Println(TweetStatusDescription)

	// Calculating end time for the bot
	endTime := time.Now().Local()
	endttimeStampString := endTime.Format("2006-01-02 15:04:05")
	unixNano = endTime.UnixNano()
	umillisec1 := unixNano / 1000000

	/*
	   Find differece betweeen start time and end time of the bot server

	*/
	diff_milli_seconds := umillisec1 - umillisec

	fmt.Println()
	fmt.Println()
	fmt.Println("+++++++++++++++++++++++++++++++++++++++")
	fmt.Println("Closing Twitter-Helper-Bot......", endttimeStampString)
	fmt.Println("Millisecond consumed ......", diff_milli_seconds)
	fmt.Println("+++++++++++++++++++++++++++++++++++++++")

}
