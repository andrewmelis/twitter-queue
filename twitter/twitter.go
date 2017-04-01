package twitter

import (
	"fmt"
	"github.com/ChimeraCoder/anaconda"
)

// these are cavs-specific keys right now
func NewTwitter() *anaconda.TwitterApi {
	anaconda.SetConsumerKey("fake") // TODO get from env
	anaconda.SetConsumerSecret("fake")
	return anaconda.NewTwitterApi("fake", "fake")
}

var t *anaconda.TwitterApi

func Tweet(s string) {
	if t == nil {
		fmt.Println("initializing new twitter api")
		t = NewTwitter()
	}
	fmt.Println(s)
	_, err := t.PostTweet(s, nil)
	if err != nil {
		fmt.Printf("error posting tweet: %s\n", err)
	}
}
