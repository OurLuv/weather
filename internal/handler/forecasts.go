package handler

import (
	"context"
	"encoding/json"
	"net/http"
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
