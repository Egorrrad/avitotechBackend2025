package handlers

import (
	"encoding/json"
	"github.com/Egorrrad/avitotechBackend2025/internal/dto"
	pkg "github.com/Egorrrad/avitotechBackend2025/internal/pkg/http"
	"net/http"
)

// CreatePVZ Создание ПВЗ (только для модераторов)
func (h *Handler) AddPVZ(w http.ResponseWriter, r *http.Request) {
	var req dto.PostPvzJSONRequestBody

	// надо проверить город и вернуть ошибку, если он не из заданного списка
	/*
		ПВЗ возможно только в трёх городах: Москва, Санкт-Петербург и Казань.
		В других городах ПВЗ завести на первых порах нельзя, в таком случае необходимо вернуть ошибку.
	*/
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		pkg.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	pvz, err := h.repository.AddPVZ(r.Context(), req)

	if err != nil {
		pkg.SendError(w, http.StatusInternalServerError, err.Error())
	}

	pkg.RespondJSON(w, http.StatusCreated, pvz)
}
