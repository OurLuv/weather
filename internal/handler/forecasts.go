package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/OurLuv/weather/internal/model"
	"github.com/gorilla/mux"
)

func (h *Handler) GetCityList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// getting all cities
	cities, err := h.service.GetCityList(context.Background())
	if err != nil {
		h.log.Error("can't get city list", "err", err.Error())
		SendError(w, "can't get city list", http.StatusInternalServerError)
		return
	}

	// sending response
	json.NewEncoder(w).Encode(cities)
}

func (h *Handler) GetShortForecast(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// getting id
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.log.Error("id is invalid", "err", err.Error())
		SendError(w, "id is invalid", http.StatusBadRequest)
		return
	}

	// getting forecast
	f, err := h.service.GetShortForecast(context.Background(), id)
	if err != nil {
		h.log.Error("can't get a forecast", "err", err.Error())
		SendError(w, "can't get a forecast", http.StatusInternalServerError)
		return
	}

	// sending response
	var shortForecast model.ShortForecast
	shortForecast.CityName = f.City.Name
	shortForecast.Country = f.City.Country
	shortForecast.Temperature = f.Temperature
	shortForecast.DateList = f.DateList
	json.NewEncoder(w).Encode(shortForecast)
}

func (h *Handler) GetDetailedForecast(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// getting id
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.log.Error("id is invalid", "err", err.Error())
		SendError(w, "id is invalid", http.StatusBadRequest)
		return
	}
	// getting time
	tStr := r.URL.Query().Get("time")

	// parsing time
	t, err := time.Parse(time.DateTime, tStr)
	if err != nil {
		h.log.Error("time is invalid", "err", err.Error())
		SendError(w, "time is invalid", http.StatusBadRequest)
		return
	}
	dt := t.Unix()

	// getting detailed forecast
	f, err := h.service.GetDetailedForecast(context.Background(), id, int(dt))
	if err != nil {
		h.log.Error("can't get a forecast", "err", err.Error())
		SendError(w, "can't get a forecast", http.StatusInternalServerError)
		return
	}

	// sending response
	json.NewEncoder(w).Encode(f)
}
