package main

import (
	"net/http"
	"os"
	"time"

	"github.com/binarydud/covidapi/router"
	"github.com/rs/zerolog"
)

func main() {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	log := zerolog.New(output).With().
		Timestamp().
		Str("role", "dev").
		Logger()
	log.Info().Msg("api gateway")
	r := router.NewRouter(log)

	http.ListenAndServe(":5000", r)

}
