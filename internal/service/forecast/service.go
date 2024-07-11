package forecast

import (
	"log/slog"

	"github.com/OurLuv/weather/internal/storage"
)

type ForecastService struct {
	Openweathermap
}

func NewForecastService(repo storage.OpenweathermapStorage, log *slog.Logger) *ForecastService {
	return &ForecastService{
		Openweathermap: NewOpenweathermap(repo, log),
	}
}
