package apiservice

import (
	"context"

	"github.com/OurLuv/weather/internal/model"
	"github.com/OurLuv/weather/internal/storage/postgres"
)

type ForecastService interface {
	GetCityList(ctx context.Context) ([]model.City, error)
}

type ForecastServiceImpl struct {
	repo postgres.ForecastStorage
}

func (s *ForecastServiceImpl) GetCityList(ctx context.Context) ([]model.City, error) {
	return s.repo.GetCityList(ctx)
}

func NewAPIService(repo postgres.ForecastStorage) *ForecastServiceImpl {
	return &ForecastServiceImpl{
		repo: repo,
	}
}
