package postgres

import (
	"context"
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
		"JOIN forecasts f ON c.id = f.city_id WHERE f.dt < $1 " +
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

func NewForcastsRepository(pool *pgxpool.Pool) *ForecastRepository {
	return &ForecastRepository{
		pool: pool,
	}
}
