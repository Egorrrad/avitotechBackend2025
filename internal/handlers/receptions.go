package handlers

import (
	"avitotechBackend2025/internal/dto"
	h "avitotechBackend2025/internal/pkg/http"
	"encoding/json"
	"github.com/oapi-codegen/runtime/types"
	"net/http"
	"time"
)

// Receptions Создание новой приемки товаров (только для сотрудников ПВЗ)
func Receptions(w http.ResponseWriter, r *http.Request) {
	var req dto.PostReceptionsJSONRequestBody

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	response := dto.Reception{
		DateTime: time.Time{},
		Id:       nil,
		PvzId:    types.UUID{},
		Status:   "",
	}

	// Возврат ответа
	h.RespondJSON(w, http.StatusCreated, response)

}
