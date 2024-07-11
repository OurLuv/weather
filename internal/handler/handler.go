package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	apiservice "github.com/OurLuv/weather/internal/service/api-service"
	"github.com/gorilla/mux"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type Handler struct {
	service apiservice.ForecastService
	log     *slog.Logger
}

func (h *Handler) InitRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/city/list", h.GetCityList).Methods("GET")
	r.HandleFunc("/city/{id}", h.GetDetailedForecast).Methods("GET")
	r.HandleFunc("/city/{id}/short-forecast", h.GetShortForecast).Methods("GET")

	h.log.Info("Init routes")
	return r
}

func SendError(w http.ResponseWriter, errorStr string, code int) {
	w.WriteHeader(code)
	response := ErrorResponse{
		Error: errorStr,
	}
	json.NewEncoder(w).Encode(response)
}

func NewHandler(service apiservice.ForecastService, log *slog.Logger) *Handler {
	return &Handler{
		service: service,
		log:     log,
	}
}
