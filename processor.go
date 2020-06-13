package main

import (
	"os"

	"github.com/binarydud/covidapi/client"
	"github.com/rs/zerolog"
)

func main() {
	logger := zerolog.New(os.Stdout).With().
		Timestamp().
		Str("role", "data processor").
		Logger()
	http := client.NewClient()
	logger.Info().Msg("calling national client")
	items, err := http.ByNational()
	if err != nil {

		logger.Fatal().Err(err).Msg("oops")
	}
	for _, item := range items {
		logger.Info().Float64("postiveAvg", item.PositiveAvg).Int("date", item.Date).Msg(item.Hash)
	}

	logger.Info().Msg("calling state client")
	states, err := http.ByStates()
	if err != nil {
		logger.Fatal().Err(err).Msg("oops")
	}
	for _, item := range states {
		logger.Info().Float64("postiveAvg", item.PositiveAvg).Int("date", item.Date).Str("state", item.State).Msg(item.Hash)
	}

}
