package twitterhelper

import (
	"github.com/dghubble/go-twitter/twitter"
)

func PostTweet(content string, client *twitter.Client) {

	// send tweet

	tweet, resp, err := client.Statuses.Update(content, nil)

	if err != nil {
		log.Println(err)
	}

	log.Printf("%+v\n", resp)
	log.Printf("%+v\n", tweet)

}
