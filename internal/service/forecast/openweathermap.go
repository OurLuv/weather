package forecast

import (
	"context"
	"encoding/json"

	//"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/OurLuv/weather/internal/model"
	"github.com/OurLuv/weather/internal/storage"
)

var (
	URLPattern = "https://api.openweathermap.org/data/2.5/forecast?lat=%f&lon=%f&appid=%s&units=metric"
)

type WeatherResponse struct {
	List []Weather `json:"list"`
}
type Weather struct {
	Dt   int  `json:"dt"`
	Main Main `json:"main"`
}
type Main struct {
	Temp float64 `json:"temp"`
}

type Openweathermap interface {
	GetForecast(ctx context.Context, key string) ([]model.Forecast, error)
	SetForecast(ctx context.Context, forecasts []model.Forecast) error
	InitService(ctx context.Context, key string)
}

type OpenweathermapImpl struct {
	repo storage.OpenweathermapStorage
	log  *slog.Logger
}

// * Getting forecasts from Openwwathermap API
func (o *OpenweathermapImpl) GetForecast(ctx context.Context, key string) ([]model.Forecast, error) {
	// getting all cities
	cities, err := o.repo.GetAllCities(ctx)
	if err != nil {
		return nil, fmt.Errorf("[Openweathermap] can't get data from storage: %w", err)
	}

	// getting forecasts
	var forecasts []model.Forecast
	var f model.Forecast
	for _, c := range cities {
		currentReq := fmt.Sprintf(URLPattern, c.Lat, c.Lon, key)
		o.log.Debug("Current request", "req", currentReq)
		r, err := http.Get(currentReq)
		if err != nil {
			return nil, fmt.Errorf("[Openweathermap] can't get data from Openweathermap API: %w", err)
		}
		// setting recieved data
		var resp WeatherResponse
		data, err := io.ReadAll(r.Body)
		if err != nil {
			return nil, err
		}
		//f.JSONStr = string(data)
		if err := json.Unmarshal(data, &resp); err != nil {
			return nil, err
		}
		f.CityId = c.Id
		f.JSONStr = string(data)
		for _, v := range resp.List {
			f.DateInt = v.Dt
			f.Temperature = v.Main.Temp
			forecasts = append(forecasts, f)
		}
	}
	o.log.Debug("Recieved all data from API")
	return forecasts, nil
}

// * passing recieved data to storage
func (o *OpenweathermapImpl) SetForecast(ctx context.Context, forecasts []model.Forecast) error {
	return o.repo.SetForecast(ctx, forecasts)
}

func (o *OpenweathermapImpl) InitService(ctx context.Context, key string) {
	forecasts, err := o.GetForecast(ctx, key)
	if err != nil {
		o.log.Error("can't get data from API", "err", err.Error())
		return
	}
	if err := o.SetForecast(ctx, forecasts); err != nil {
		o.log.Error("can't set data", "err", err.Error())
		return
	}

	o.log.Info("Forecast is updated")
}

func NewOpenweathermap(repo storage.OpenweathermapStorage, log *slog.Logger) *OpenweathermapImpl {
	return &OpenweathermapImpl{
		repo: repo,
		log:  log,
	}
}
