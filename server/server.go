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

	port := ":5000"
	log.Info().Str("port", port).Msg("api gateway")
	r := router.NewRouter(log)

	http.ListenAndServe(port, r)

}
