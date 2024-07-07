package postgres

import (
	"context"

	"github.com/OurLuv/weather/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OpenweathermapRepository struct {
	pool *pgxpool.Pool
}

func (r *OpenweathermapRepository) GetAllCities(ctx context.Context) ([]model.City, error) {
	query := "SELECT * FROM cities"
	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	var cities []model.City
	var c model.City
	for rows.Next() {
		if err := rows.Scan(&c.Id, &c.Name, &c.Country, &c.Lat, &c.Lon); err != nil {
			return nil, err
		}
		cities = append(cities, c)
	}

	return cities, nil
}

func (r *OpenweathermapRepository) SetForecast(ctx context.Context, cities []model.Forecast) error {
	return nil
}

func NewOpenweathermapRepository(pool *pgxpool.Pool) *OpenweathermapRepository {
	return &OpenweathermapRepository{
		pool: pool,
	}
}
