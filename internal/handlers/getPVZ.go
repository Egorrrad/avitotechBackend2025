package handlers

import (
	"avitotechBackend2025/internal/dto"
	h "avitotechBackend2025/internal/pkg/http"
	"github.com/gorilla/schema"
	"net/http"
)

func GetPVZ(w http.ResponseWriter, r *http.Request) {
	var req dto.GetPvzParams

	// нужно еще добавить валидацию для диапазонов
	err := schema.NewDecoder().Decode(&req, r.URL.Query())
	if err != nil {
		h.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	h.RespondJSON(w, http.StatusOK, req)
}
