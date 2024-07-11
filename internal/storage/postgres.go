package storage

import (
	"context"
	"fmt"

	"github.com/OurLuv/weather/internal/config"
	"github.com/OurLuv/weather/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OpenweathermapStorage interface {
	GetAllCities(ctx context.Context) ([]model.City, error)
	SetForecast(ctx context.Context, forecast []model.Forecast) error
}

type ForecastStorage interface {
	GetCityList(ctx context.Context) ([]model.City, error)
	GetShortForecast(ctx context.Context, cityId int) (*model.Forecast, error)
	GetDetailedForecast(ctx context.Context, cityId int, dt int) (string, error)
}

func NewPostgresPool(ctx context.Context, cfg config.Config) (*pgxpool.Pool, error) {
	coonStr := fmt.Sprintf("postgres://%s:%s@localhost:5432/%s", cfg.User, cfg.Password, cfg.DatabaseName)
	pool, err := pgxpool.New(ctx, coonStr)
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, err
	}

	return pool, nil
}
