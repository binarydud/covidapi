package main

import (
	"context"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/binarydud/covidapi/client"
	"github.com/rs/zerolog"
)

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, event MyEvent) {
	log := zerolog.New(os.Stdout).With().
		Timestamp().
		Str("role", "data processor").
		Logger()
	client := client.NewClient()
	log.Info().Msg("Processing state logs")
	client.ByStates()
	// get state historical data. Calculate averages for each state
	client.ByNational()
	// get national historical data. Calculate avergages. save to dynamo
}
func main() {
	lambda.Start(HandleRequest)
}
