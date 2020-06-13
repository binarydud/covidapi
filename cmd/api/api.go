package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/binarydud/covidapi/router"
	"github.com/binarydud/pylon"
	"github.com/rs/zerolog"
)

func main() {
	fmt.Println(os.Args)
	log := zerolog.New(os.Stdout).With().
		Timestamp().
		Logger()
	r := router.NewRouter(log)
	lambda.Start(pylon.GatewayProxyEvent(r))
}
