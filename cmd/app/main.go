package main

import (
	"mini-stats-server/config"
	"mini-stats-server/internal/repository"
	"mini-stats-server/internal/server"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	log.Info().Msg("starting mini stats server")

	conf := config.New()

	log.Info().Interface("dict", conf).Msg("configuration")

	repo := repository.New(conf)

	server := server.New(repo)
	server.Start()
}
