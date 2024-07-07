package forecast

import (
	"context"

	"github.com/OurLuv/weather/internal/model"
)

type ForecastService interface {
	GetForecast(cfg context.Context, key string) ([]model.Forecast, error)
	SetForecast(cfg context.Context, forecasts []model.Forecast) error
}
