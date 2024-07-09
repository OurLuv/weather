package main

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/OurLuv/weather/internal/config"
	"github.com/OurLuv/weather/internal/handler"
	apiservice "github.com/OurLuv/weather/internal/service/api-service"
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
	repo1 := postgres.NewOpenweathermapRepository(pool)
	repo2 := postgres.NewForcastsRepository(pool)

	// init services
	//openweathermap
	log.Info("Init Openweathermap service")
	var forecastService forecast.ForecastService
	forecastService = forecast.NewOpenweathermap(repo1, log)
	forecasts, err := forecastService.GetForecast(context.Background(), cfg.KEY)
	if err != nil {
		log.Error("can't get data from API", slog.String("err", err.Error()))
	}
	if err := forecastService.SetForecast(context.Background(), forecasts); err != nil {
		log.Error("can't get data from API", slog.String("err", err.Error()))
	}
	log.Info("Set new forecasts")
	//api's service
	apiService := apiservice.NewAPIService(repo2)

	// init server
	h := handler.NewHandler(apiService, log)
	router := h.InitRoutes()
	server := handler.NewServer(router)
	log.Info("Server is started")
	if err := server.Start(); err != nil {
		log.Error("can't start server", "err", err.Error())
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
