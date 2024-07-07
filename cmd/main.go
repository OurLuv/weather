package main

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/OurLuv/weather/internal/config"
	"github.com/OurLuv/weather/internal/service/forecast"
	"github.com/OurLuv/weather/internal/storage/postgres"
	log "github.com/charmbracelet/log"
)

func main() {

	// config
	cfg := config.MustLoad()

	// init logger
	log := InitLogger()
	log.Info("Starting application")
	log.Debug("Starting with config", slog.Any("cfg", cfg))

	// init repo
	log.Info("Init repo")
	pool, err := postgres.NewPostgresPool(context.Background(), *cfg)
	if err != nil {
		log.Error("can't init connection to db", slog.String("err", err.Error()))
		os.Exit(1)
	}
	repo := postgres.NewOpenweathermapRepository(pool)

	// init Openweathermap service
	log.Info("Init Openweathermap service")
	var forecastService forecast.ForecastService
	forecastService = forecast.NewOpenweathermap(repo, log)
	_, err = forecastService.GetForecast(context.Background(), cfg.KEY)
	if err != nil {
		log.Error("can't get data from API", slog.String("err", err.Error()))
	}
}

func InitLogger() *slog.Logger {
	handler := log.NewWithOptions(os.Stdout, log.Options{
		Level:           -4,
		ReportTimestamp: true,
		ReportCaller:    true,
		CallerOffset:    0,
		TimeFormat:      time.Kitchen,
		Prefix:          "🍪 ",
	})
	logger := slog.New(handler)

	return logger
}
