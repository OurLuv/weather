package apiservice

import (
	"context"
	"encoding/json"

	"github.com/OurLuv/weather/internal/model"
	"github.com/OurLuv/weather/internal/storage/postgres"
)

type ForecastService interface {
	GetCityList(ctx context.Context) ([]model.City, error)
	GetShortForecast(ctx context.Context, cityId int) (*model.Forecast, error)
	GetDetailedForecast(ctx context.Context, cityId int, dt int) (*model.WeatherData, error)
}

type ForecastServiceImpl struct {
	repo postgres.ForecastStorage
}

func (s *ForecastServiceImpl) GetCityList(ctx context.Context) ([]model.City, error) {
	return s.repo.GetCityList(ctx)
}

func (s *ForecastServiceImpl) GetShortForecast(ctx context.Context, cityId int) (*model.Forecast, error) {
	return s.repo.GetShortForecast(ctx, cityId)
}

func (s *ForecastServiceImpl) GetDetailedForecast(ctx context.Context, cityId int, dt int) (*model.WeatherData, error) {

	// getting data from database
	jsonStr, err := s.repo.GetDetailedForecast(ctx, cityId, dt)
	if err != nil {
		return nil, err
	}

	// unmarshaling
	var weather model.WeatherData
	if err := json.Unmarshal([]byte(jsonStr), &weather); err != nil {
		return nil, err
	}

	// getting only needed
	for i := range weather.List {
		if weather.List[i].Dt == dt {
			weather.Weather = weather.List[i]
			break
		}
	}
	weather.List = []model.Weather{}

	return &weather, nil

}

func NewAPIService(repo postgres.ForecastStorage) *ForecastServiceImpl {
	return &ForecastServiceImpl{
		repo: repo,
	}
}
