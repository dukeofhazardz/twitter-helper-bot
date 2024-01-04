package twitterhelper

import (
	"fmt"
	"os"
	"strconv"

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

func RetweetingTweet(TweetStatusIds []string, client *twitter.Client) string {
	var retweet_status_description string = "Retweeted  unsuccessfull"
	var RetweetStatus twitter.StatusRetweetParams
	var tweetcount int = 0

	fmt.Println("======================================================")
	fmt.Println("Started Reweeting Hashtag.....")
	for itr := range TweetStatusIds {

		if len(TweetStatusIds[itr]) > 0 {
			ID, _ := strconv.ParseInt(TweetStatusIds[itr], 10, 64)

			retweet, _, err := client.Statuses.Retweet(ID, &RetweetStatus)

			if err != nil {
				log.Error("Error in  Reweeting")
			} else {
				tweetcount = tweetcount + 1
				log.Infoln("Successfully reweeted", retweet.IDStr)
				retweet_status_description = "Successfully reweeted"
			}
		}

	}

	fmt.Println("Successfully  Reweeted", tweetcount, "Hashtag.....")
	fmt.Println("======================================================")
	return retweet_status_description
}
