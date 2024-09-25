package main

import (
	"embed"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/templhtmx/handler"
	"github.com/templhtmx/pkg/supabase"
)

//go:embed public
var FS embed.FS

func main() {

	if err := initServer(); err != nil {
		log.Fatal().AnErr("error", err).Msg("Server Initialization Failed")
	}

	port := os.Getenv("HTTP_LISTEN_ADDR")
	router := chi.NewMux()

	router.Use(handler.WithUser)

	router.Handle("/*", http.StripPrefix("/", http.FileServer(http.FS(FS))))
	router.Get("/", handler.Make(handler.HandleHomeIndex))
	router.Get("/signin", handler.Make(handler.HandleSigninIndex))
	router.Post("/signin", handler.Make(handler.HandleSigninCreate))

	log.Info().Str("port", port).Msg("Application running")
	log.Fatal().AnErr("error", http.ListenAndServe(port, router)).Msg("ListenAndServe Failed")
}

func initServer() error {

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	if err := godotenv.Load(); err != nil {
		return err
	}
	return supabase.Init()
}
