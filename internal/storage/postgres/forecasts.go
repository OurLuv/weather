package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/OurLuv/weather/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ForecastRepository struct {
	pool *pgxpool.Pool
}

func (r *ForecastRepository) GetCityList(ctx context.Context) ([]model.City, error) {
	var cities []model.City
	var c model.City

	query := "SELECT c.id, c.name, c.country, c.lat, c.lon FROM cities c " +
		"JOIN forecasts f ON c.id = f.city_id WHERE f.dt > $1 " +
		"GROUP BY c.id " +
		"ORDER BY c.name"
	rows, err := r.pool.Query(ctx, query, time.Now().Unix())
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		if err := rows.Scan(&c.Id, &c.Name, &c.Country, &c.Lat, &c.Lon); err != nil {
			return nil, err
		}
		cities = append(cities, c)
	}

	return cities, nil
}

func (r *ForecastRepository) GetShortForecast(ctx context.Context, cityId int) (*model.Forecast, error) {
	var f model.Forecast

	query := "SELECT c.id, c.name, c.country, f.temp, f.dt FROM cities c " +
		"JOIN forecasts f ON c.id = f.city_id WHERE f.dt > $1 AND c.id = $2 " +
		"ORDER BY c.name"

	rows, err := r.pool.Query(ctx, query, time.Now().Unix(), cityId)
	if err != nil {
		return nil, err
	}
	i := 0.0
	tempAv := 0.0
	for rows.Next() {
		i++
		if err := rows.Scan(&f.City.Id, &f.City.Name, &f.City.Country, &f.Temperature, &f.DateInt); err != nil {
			return nil, err
		}

		f.DateList = append(f.DateList, time.Unix(int64(f.DateInt), 0).Format(time.DateTime))
		tempAv += f.Temperature
	}
	if i == 0 {
		return nil, fmt.Errorf("no rows in result set")
	}
	f.Temperature = tempAv / i

	return &f, nil
}

func (r *ForecastRepository) GetDetailedForecast(ctx context.Context, cityId int, dt int) (string, error) {
	var jsonStr string
	query := "SELECT json FROM forecasts WHERE city_id=$1 AND dt = $2"
	if err := r.pool.QueryRow(ctx, query, cityId, dt).Scan(&jsonStr); err != nil {
		return "", err
	}
	return jsonStr, nil
}

func NewForcastsRepository(pool *pgxpool.Pool) *ForecastRepository {
	return &ForecastRepository{
		pool: pool,
	}
}
