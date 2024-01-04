package twitterhelper

import (
	"fmt"
	"os"

	"github.com/dghubble/go-twitter/twitter"
)

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

func FilterTweetstatusId(SearchQueries twitter.Search, client *twitter.Client) []string {

	fmt.Printf("======================================================")
	fmt.Println("Started retrieving status ID......")

	TweetStatusIds := make([]string, 10)
	for ids := range SearchQueries.Statuses {
		//fmt.Println(a.Statuses[ids].IDStr)

		TweetStatusIds = append(TweetStatusIds, SearchQueries.Statuses[ids].IDStr)

	}

	fmt.Println("Completed retrieving", len(TweetStatusIds), "status ID")
	fmt.Printf("======================================================")

	return TweetStatusIds
}
