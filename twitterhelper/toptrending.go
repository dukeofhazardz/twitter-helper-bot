package twitterhelper

import (
	"encoding/json"
	"fmt"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func TopTrending(client *twitter.Client) []string {
	fmt.Printf("======================================================")
	fmt.Println("Fetching top trending hashtags lists from the United States......")

	TrendListName := make([]string, 5, 5)
	// trending lists of the United States
	trendlist, _, err := client.Trends.Place(int64(23424977), &twitter.TrendsPlaceParams{WOEID: 23424977})
	trendJSON, _ := json.MarshalIndent(trendlist, "", " ")
	//fmt.Println(string(trendJSON))

	// Unmarshaling JSON DATA
	err = json.Unmarshal(trendJSON, &trendlist)
	if err != nil {
		log.Error("Error in Unmarshaling")
	}

	//fmt.Println(trendlist[0].Trends[0].Name)

	for Ids := range trendlist[0].Trends {
		TrendListName = append(TrendListName, trendlist[0].Trends[Ids].Name)
	}

	fmt.Printf("======================================================")
	fmt.Println("Completed retrieving", len(TrendListName), "top trending hashtags lists from the United States")

	return TrendListName
}
