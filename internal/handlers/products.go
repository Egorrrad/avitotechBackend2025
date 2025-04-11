package handlers

import (
	"avitotechBackend2025/internal/dto"
	h "avitotechBackend2025/internal/pkg/http"
	"encoding/json"
	"github.com/oapi-codegen/runtime/types"
	"net/http"
)

// Products Добавление товара в текущую приемку (только для сотрудников ПВЗ)
func Products(w http.ResponseWriter, r *http.Request) {
	var req dto.PostProductsJSONRequestBody

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	response := dto.Product{
		DateTime:    nil,
		Id:          nil,
		ReceptionId: types.UUID{},
		Type:        "",
	}

	h.RespondJSON(w, http.StatusCreated, response)
}
