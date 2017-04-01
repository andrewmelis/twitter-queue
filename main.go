package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/andrewmelis/twitter-queue/queue"
	"github.com/andrewmelis/twitter-queue/twitter"
)

var q queue.Queue

func main() {
	http.HandleFunc("/tweet", tweetHandler)

	q = queue.NewQueue(twitter.Tweet)
	// q = queue.NewQueue(func(s string) {
	// fmt.Println(s)
	// })

	log.Fatal(http.ListenAndServe(":8083", nil))
}

func tweetHandler(w http.ResponseWriter, r *http.Request) {
	var pbp PlayByPlayGame
	dec := json.NewDecoder(r.Body)
	for dec.More() {

		err := dec.Decode(&pbp)
		if err != nil {
			log.Printf("error decoding game: %s\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, `{"error":"server error occurred"}`)
			return
		}
	}

	if strings.Contains(pbp.Game.GameCode(), "CLE") { // HACK -- need additional creds
		// extract
		for _, p := range pbp.Plays {
			err := q.Enqueue(p.Description) // includes score + stats -- TODO
			if err != nil {
				log.Printf("error enqueueing play: %+v\n", p)
			}
		}
	}
	w.WriteHeader(http.StatusAccepted)
}

type PlayByPlayGame struct {
	Game
	Plays []Play
}

type Game struct {
	Id           string    `json:"gameId"`
	StartTime    time.Time `json:"startTimeUTC"`
	VisitingTeam Team      `json:"vTeam"`
	HomeTeam     Team      `json:"hTeam"`
	Period       Period    `json:"period"`
	Active       bool      `json:"isGameActivated"`
}

func (g Game) GameCode() string {
	return fmt.Sprintf("%s%s", g.VisitingTeam.TriCode, g.HomeTeam.TriCode)
}

// func (p Play) String() string {
// 	return fmt.Sprintf("%s\n[%s - %s]",Play.Formatted.Description.
// }

type Play struct {
	Clock            string        `json:"clock"`
	Description      string        `json:"description"`
	PersonId         string        `json:"personId"`
	TeamId           string        `json:"teamId"`
	VistingTeamScore string        `json:"vTeamScore"`
	HomeTeamScore    string        `json:"hTeamScore"`
	IsScoreChange    bool          `json:"isScoreChange"`
	Formatted        FormattedPlay `json:"formatted"`
}

type Team struct {
	Id      string `json:"teamId"`
	TriCode string `json:"triCode"`
}

type Period struct {
	Current int
}

type FormattedPlay struct {
	Description string `json:"description"`
}
