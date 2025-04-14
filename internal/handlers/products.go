package handlers

import (
	"encoding/json"
	"github.com/Egorrrad/avitotechBackend2025/internal/dto"
	pkg "github.com/Egorrrad/avitotechBackend2025/internal/pkg/http"
	"net/http"
)

// Products Добавление товара в текущую приемку (только для сотрудников ПВЗ)
func (h *Handler) Products(w http.ResponseWriter, r *http.Request) {
	var req dto.PostProductsJSONRequestBody

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		pkg.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	product, err := h.repository.AddProduct(r.Context(), req)

	if err != nil {
		pkg.SendError(w, http.StatusInternalServerError, err.Error())
	}

	pkg.RespondJSON(w, http.StatusCreated, product)
}
