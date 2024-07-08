package forecast

import (
	"context"

	"github.com/OurLuv/weather/internal/model"
)

type ForecastService interface {
	GetForecast(ctx context.Context, key string) ([]model.Forecast, error)
	SetForecast(ctx context.Context, forecasts []model.Forecast) error
}
