package main

import (
	"errors"
	"github.com/antlko/go-currency-convertor/internal"
	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
	"log/slog"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil && !errors.Is(err, os.ErrNotExist) {
		slog.Error("environment variables not loaded", "error_message", err.Error())
		os.Exit(1)
	}

	var config internal.AppConfig
	if err := env.Parse(&config); err != nil {
		slog.Error("environment variabled not parsed", "error_message", err.Error())
		os.Exit(1)
	}

	if err := internal.InitService(config); err != nil {
		slog.Error("Service error", "error", err)
	}
}
