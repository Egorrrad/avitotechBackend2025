package handlers

import (
	"avitotechBackend2025/internal/dto"
	h "avitotechBackend2025/internal/pkg/http"
	"encoding/json"
	"net/http"
)

// CreatePVZ Создание ПВЗ (только для модераторов)
func CreatePVZ(w http.ResponseWriter, r *http.Request) {
	var req dto.PostPvzJSONRequestBody

	// Проверка прав доступа и создание ПВЗ

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Логика создания ПВЗ

	pvz := dto.PVZ{
		City:             "",
		Id:               nil,
		RegistrationDate: nil,
	}

	h.RespondJSON(w, http.StatusCreated, pvz)
}
