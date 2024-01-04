package twitterhelper

import (
	"encoding/json"
	"fmt"

	"github.com/dghubble/go-twitter/twitter"
)

func Searchqueries(TrendListName []string, client *twitter.Client) twitter.Search {
	var SearchQueries twitter.Search
	for itr := range TrendListName {
		if len(TrendListName[itr]) > 0 {

			searchedtweet, _, err := client.Search.Tweets(&twitter.SearchTweetParams{

				Query: TrendListName[itr],
			})

			// search Query

			if err != nil {
				log.Error("Failed to fetch search query", TrendListName[itr])
				log.Error(err)
			} else {
				fmt.Println("======================================================")
				fmt.Println("successfully retrieved the hastag", TrendListName[itr])
				fmt.Println("======================================================")
			}

			jsonData, err := json.MarshalIndent(searchedtweet, "", " ")
			if err != nil {
				log.Error("Failed to marshal searchedtweet")
			}

			// print JSON DATA from search Query
			//log.Printf("%+v\n", string(jData))

			// Unmarshaling JSON DATA
			err = json.Unmarshal(jsonData, &SearchQueries)
			if err != nil {
				log.Error("Error in Unmarshaling")
			}

		}
	}

	return SearchQueries
}
