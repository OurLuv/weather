package model

import "time"

type Forecast struct {
	CityId      int       `json:"city_id,omitempty"`
	City        City      `json:"city,omitempty"`
	Temperature float64   `json:"temp"`
	Date        time.Time `json:"date,omitempty"`
	DateList    []string  `json:"date_list,omitempty"`
	DateInt     int       `json:"dt,omitempty"`
	JSONStr     string    `json:"json_str"`
}

type ShortForecast struct {
	CityName    string   `json:"city,omitempty"`
	Country     string   `json:"country"`
	Temperature float64  `json:"temp"`
	DateList    []string `json:"date_list,omitempty"`
}

type WeatherData struct {
	Weather Weather   `json:"forecast,omitempty"`
	List    []Weather `json:"list,omitempty"`
	City    CityOWM   `json:"city"`
}

type Weather struct {
	Dt         int           `json:"dt"`
	Main       Main          `json:"main"`
	Weather    []WeatherDesc `json:"weather"`
	Clouds     Clouds        `json:"clouds"`
	Wind       Wind          `json:"wind"`
	Visibility int           `json:"visibility"`
	Pop        float64       `json:"pop"`
	Rain       *Rain         `json:"rain,omitempty"`
	DtTxt      string        `json:"dt_txt"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int     `json:"pressure"`
	SeaLevel  int     `json:"sea_level"`
	GrndLevel int     `json:"grnd_level"`
	Humidity  int     `json:"humidity"`
	TempKf    float64 `json:"temp_kf"`
}

type WeatherDesc struct {
	Main        string `json:"main"`
	Description string `json:"description"`
}

type Clouds struct {
	All int `json:"all"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int     `json:"deg"`
	Gust  float64 `json:"gust"`
}

type Rain struct {
	ThreeHours float64 `json:"3h"`
}

type CityOWM struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Coord      Coord  `json:"coord"`
	Country    string `json:"country"`
	Population int    `json:"population"`
	Timezone   int    `json:"timezone"`
	Sunrise    int    `json:"sunrise"`
	Sunset     int    `json:"sunset"`
}

type Coord struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}
