package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/binarydud/covidapi/client"
	"github.com/rs/zerolog/hlog"
	"github.com/slack-go/slack"
)

// AuthorizeHandler ...
func AuthorizeHandler(w http.ResponseWriter, r *http.Request) {
	log := hlog.FromRequest(r)
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	clientID := os.Getenv("CLIENTID")
	clientSecret := os.Getenv("CLIENTSECRET")
	code := r.URL.Query().Get("code")

	_, _, err := slack.GetOAuthToken(client, clientID, clientSecret, code, "")
	if err != nil {
		log.Err(err).Msg("error authorizing")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Info().Msg("Authorization request")
	w.Write([]byte("success, app installed"))
}

// CommandHandler ...
func CommandHandler(w http.ResponseWriter, r *http.Request) {
	s, err := slack.SlashCommandParse(r)
	log := hlog.FromRequest(r)
	if err != nil {
		log.Error().Err(err).Msg("error parsing command")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error parsing command"))
		return
	}
	log.Info().
		Str("user", s.UserID).
		Str("status", "ok").
		Msg("parsed command")
	state := s.Text
	client := client.New()
	item, err := client.ByState(state)
	if err != nil {
		log.Err(err).Msg("error calling api")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// loc, _ := time.LoadLocation("America/Chicago")
	attachments := []slack.Attachment{}

	attachments = append(attachments, slack.Attachment{
		Text: fmt.Sprintf("7 day averages for %s", item.State),
	})
	attachments = append(attachments, slack.Attachment{
		Text: fmt.Sprintf("Average Postive Case Count %f", item.PositiveAvg),
	})
	attachments = append(attachments, slack.Attachment{
		Text: fmt.Sprintf("Average Daily Fatality Count %f", item.DeathsAvg),
	})
	attachments = append(attachments, slack.Attachment{
		Text: fmt.Sprintf("Average Percentage of positive tests %f", item.PercentagePositive),
	})
	attachments = append(attachments, slack.Attachment{
		Text: fmt.Sprintf("Most recent day's positive tests %d", item.PositiveIncrease),
	})
	attachments = append(attachments, slack.Attachment{
		Text: fmt.Sprintf("Most recent day's fatalities %d", item.DeathIncrease),
	})
	newPositiveCases := *item.PositiveIncrease
	newTests := *item.TotalTestResultsIncrease
	attachments = append(attachments, slack.Attachment{
		Text: fmt.Sprintf("Most recent day's percentage of tests that are positive %f", float64(newPositiveCases/newTests)),
	})

	/*
		for _, state := range items {
			hospitalized := (float64(state.Hospitalized) / float64(state.Positive)) * 100
			attachment := slack.Attachment{
				Text: fmt.Sprintf("%d positive cases, %% hospitalized %f in %s, last checked %s", state.Positive, hospitalized, state.Province, state.DateChecked.In(loc).Format("Mon, 2 Jan 2006 15:04:05 MST")),
			}
			attachments = append(attachments, attachment)
		}
	*/
	message := &slack.Msg{ResponseType: slack.ResponseTypeInChannel, Attachments: attachments, Text: "Covid stats"}

	body, err := json.Marshal(message)
	if err != nil {
		log.Err(err).Msg("error creating message")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// response := fmt.Sprintf("Hello %s", s.Text)
	w.Header().Add("Content-Type", "application/json")
	w.Write(body)
}
