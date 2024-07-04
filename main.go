package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	router := chi.NewMux()

	log.Fatal().AnErr("error", http.ListenAndServe(os.Getenv("HTTP_LISTEN_ADDR"), router)).Msg("ListenAndServe Failed")
}

func initServer() error {

	return nil

}
