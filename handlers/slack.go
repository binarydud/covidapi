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
		Str("state", s.Text).
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
	log.Info().
		Str("state", state).
		Int("PositiveCases", *item.PositiveIncrease).
		Int("Tests", *item.TotalTestResultsIncrease).
		Float64("PostiveAVG", item.PositiveAvg).
		Msg("state data")

	attachments := []slack.Attachment{}

	attachments = append(attachments, slack.Attachment{
		Text: fmt.Sprintf("7 day trailing averages for %s", item.State),
	})
	attachments = append(attachments, slack.Attachment{
		Text: fmt.Sprintf("Average Postive Case Count %f", item.PositiveAvg),
	})
	attachments = append(attachments, slack.Attachment{
		Text: fmt.Sprintf("Average Daily Fatality Count %f", item.DeathsAvg),
	})
	/*
		attachments = append(attachments, slack.Attachment{
			Text: fmt.Sprintf("Average Percentage of positive tests %f", item.PercentagePositive),
		})
	*/
	attachments = append(attachments, slack.Attachment{
		Text: "Most recent daily stats",
	})
	attachments = append(attachments, slack.Attachment{
		Text: fmt.Sprintf("Daily Positive tests %d", *item.PositiveIncrease),
	})
	attachments = append(attachments, slack.Attachment{
		Text: fmt.Sprintf("Daily tests %d", *item.TotalTestResultsIncrease),
	})
	attachments = append(attachments, slack.Attachment{
		Text: fmt.Sprintf("Daily fatalities %d", *item.DeathIncrease),
	})
	newPositiveCases := *item.PositiveIncrease
	newTests := *item.TotalTestResultsIncrease
	percent := float64(newPositiveCases) / float64(newTests) * 100
	log.Info().Float64("percentage", float64(newPositiveCases/newTests)).Msg("checking percentage")
	attachments = append(attachments, slack.Attachment{
		Text: fmt.Sprintf("Percentage of tests that are positive %f", percent),
	})
	attachments = append(attachments, slack.Attachment{
		Text: "Total stats",
	})
	attachments = append(attachments, slack.Attachment{
		Text: fmt.Sprintf("Positive Cases %d", *item.Positive),
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
	message := &slack.Msg{ResponseType: slack.ResponseTypeInChannel, Attachments: attachments, Text: fmt.Sprintf("Covid stats %s", state)}

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
