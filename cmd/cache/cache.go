package main

import (
	"context"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/binarydud/covidapi/client"
	"github.com/binarydud/covidapi/db"
	"github.com/binarydud/covidapi/types"
	"github.com/rs/zerolog"
)

func updateToday(ctx context.Context) error {
	dbclient := db.New()
	logger := zerolog.New(os.Stdout).With().
		Timestamp().
		Str("role", "covid processor").
		Logger()
	http := client.New()
	items, err := http.ByNational()
	if err != nil {

		logger.Fatal().Err(err).Msg("oops")
	}
	last := items[len(items)-1]
	err = dbclient.PutUS(last)
	if err != nil {
		logger.Fatal().Err(err).Msg("error saving todays current us data")
	}
	states, err := http.ByStates()
	if err != nil {
		logger.Fatal().Err(err).Msg("oops")
	}
	m := make(map[string][]types.State)
	for _, i := range states {
		m[i.State] = append(m[i.State], i)
	}
	for _, stateValues := range m {
		lastState := stateValues[len(stateValues)-1]
		err = dbclient.PutState(lastState)
		if err != nil {
			logger.Fatal().Err(err).Msg("error saving todays current us data")
		}
	}
	return nil
}
func handleRequest(ctx context.Context) error {
	dbclient := db.New()
	logger := zerolog.New(os.Stdout).With().
		Timestamp().
		Str("role", "covid processor").
		Logger()
	http := client.New()
	logger.Info().Msg("calling national client")
	items, err := http.ByNational()
	if err != nil {

		logger.Fatal().Err(err).Msg("oops")
	}

	for _, item := range items {
		logger.Info().Str("hash", item.Hash).Int("date", item.Date).Msg(item.Hash)
		err := dbclient.PutUS(item)
		if err != nil {
			logger.Fatal().Err(err).Msg("error saving ")
		}
	}

	logger.Info().Msg("calling state client")
	states, err := http.ByStates()
	if err != nil {
		logger.Fatal().Err(err).Msg("oops")
	}
	logger.Info().Int("number", len(states)).Msg("got some data")
	for _, item := range states {
		//logger.Debug().Int("date", item.Date).Str("state", item.State).Msg(item.Hash)
		err := dbclient.PutState(item)
		if err != nil {
			logger.Fatal().Err(err).Msg("error saving ")
		}
	}
	return nil
}
func main() {
	lambda.Start(updateToday)
	//updateToday(context.Background())
	// handleRequest(context.Background())
}
