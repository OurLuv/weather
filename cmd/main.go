package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/OurLuv/weather/internal/config"
	"github.com/OurLuv/weather/internal/handler"
	apiservice "github.com/OurLuv/weather/internal/service/api-service"
	"github.com/OurLuv/weather/internal/service/forecast"
	"github.com/OurLuv/weather/internal/storage"
	log "github.com/charmbracelet/log"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// config
	cfg := config.MustLoad()

	// init logger
	log := InitLogger()
	log.Info("Starting application")
	log.Debug("Starting with config", slog.Any("cfg", cfg))

	// init repo
	log.Info("Init repo")
	pool, err := storage.NewPostgresPool(context.Background(), *cfg)
	if err != nil {
		log.Error("can't init connection to db", slog.String("err", err.Error()))
		os.Exit(1)
	}
	repo1 := storage.NewOpenweathermapRepository(pool)
	repo2 := storage.NewForcastsRepository(pool)

	// init services
	log.Info("Init services")
	//openweathermap
	forecastService := forecast.NewForecastService(repo1, log)
	forecastService.InitService(context.Background(), cfg.KEY)
	//api's service
	apiService := apiservice.NewAPIService(repo2)

	// start server
	h := handler.NewHandler(apiService, log)
	router := h.InitRoutes()
	server := handler.NewServer(router)
	log.Info("Server is started")
	go func() {
		if err := server.Start(); err != nil {
			log.Debug("server is off", slog.String("err", err.Error()))
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-ctx.Done()
	log.Info("Shutting down")
	err = server.ShutDown()
	if err != nil {
		log.Error("Error while shutting down the server", slog.String("err", err.Error()))
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
