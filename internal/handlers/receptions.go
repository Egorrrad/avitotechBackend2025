package handlers

import (
	"encoding/json"
	"github.com/Egorrrad/avitotechBackend2025/internal/dto"
	pkg "github.com/Egorrrad/avitotechBackend2025/internal/pkg/http"
	"net/http"
)

// Receptions Создание новой приемки товаров (только для сотрудников ПВЗ)
func (h *Handler) Receptions(w http.ResponseWriter, r *http.Request) {
	var req dto.PostReceptionsJSONRequestBody

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		pkg.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	reception, err := h.Repository.AddReception(r.Context(), req)

	if err != nil {
		pkg.SendError(w, http.StatusInternalServerError, err.Error())
	}

	pkg.RespondJSON(w, http.StatusCreated, reception)

}
