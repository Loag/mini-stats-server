package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/rs/zerolog/log"
)

type Config struct {
	RepoPath  string
	RepoPort  int
	RepoToken string
}

func get_val(val string) string {
	value, ok := os.LookupEnv(val)
	if !ok {
		panic(fmt.Sprintf("required env var not set: %s", val))
	}
	return value
}

func New() *Config {
	repoPath := get_val("REPO_PATH")
	repoToken := get_val("REPO_TOKEN")
	repoPortStr := get_val("REPO_PORT")

	repoPort, err := strconv.Atoi(repoPortStr)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to get repo port")
	}

	return &Config{
		RepoPath:  repoPath,
		RepoToken: repoToken,
		RepoPort:  repoPort,
	}
}
