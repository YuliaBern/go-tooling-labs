package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/YuliaBern/lab1-tooling/internal"
)

func main() {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	config, err := internal.LoadConfig()

	if err != nil {
		logger.Error().Err(err).Msg("config error")
		return
	}

	logger.Info().
		Str("app", config.AppName).
		Int("port", config.Port).
		Msg("application started")
}