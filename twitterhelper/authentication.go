package twitterhelper

import (
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/joho/godotenv"
)

type TwitterCredentials struct {
	Consumerkey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

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

// Check client credentials
func Twitter(creds *TwitterCredentials) (*twitter.Client, error) {

	// Storing configuration value
	Consumerkey, ConsumerSecret, AccessToken, AccessTokenSecret, err := goDotEnvVariable("TWITTER_CONSUMER_KEY", "TWITTER_CONSUMER_SECRET", "TWITTER_ACCESS_TOKEN", "TWITTER_ACCESS_SECRET")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	config := oauth1.NewConfig(Consumerkey, ConsumerSecret)
	token := oauth1.NewToken(AccessToken, AccessTokenSecret)

	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}

	user, _, err := client.Accounts.VerifyCredentials(verifyParams)

	if err != nil {
		return nil, err
	}
	log.Infof("user Acconut :\n%+v\n", user)

	return client, nil
}

// use godot package to load/read the .env file and
// return the value of the keys
func goDotEnvVariable(key ...string) (string, string, string, string, error) {

	// Load .env file
	err := godotenv.Load(".env")

	Consumerkey := os.Getenv("TWITTER_CONSUMER_KEY")
	ConsumerSecret := os.Getenv("TWITTER_CONSUMER_SECRET")
	AccessToken := os.Getenv("TWITTER_ACCESS_TOKEN")
	AccessTokenSecret := os.Getenv("TWITTER_ACCESS_SECRET")

	return Consumerkey, ConsumerSecret, AccessToken, AccessTokenSecret, err

}
