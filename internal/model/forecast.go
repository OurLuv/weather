package model

import "time"

type Forecast struct {
	CityId      int       `json:"city_id"`
	Temperature float64   `json:"temp"`
	Date        time.Time `json:"date"`
	DateInt     int       `json:"dt"`
	JSONStr     string    `json:"json_str"`
}
