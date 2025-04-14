package handlers

import (
	"github.com/Egorrrad/avitotechBackend2025/internal/dto"
	pkg "github.com/Egorrrad/avitotechBackend2025/internal/pkg/http"
	"github.com/gorilla/schema"
	"net/http"
)

func (h *Handler) GetPVZ(w http.ResponseWriter, r *http.Request) {
	var req dto.GetPvzParams

	// нужно еще добавить валидацию для диапазонов
	err := schema.NewDecoder().Decode(&req, r.URL.Query())
	if err != nil {
		pkg.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	pvzArr, err := h.repository.GetAllPvz(r.Context(), req)

	if err != nil {
		pkg.SendError(w, http.StatusInternalServerError, err.Error())
	}

	pkg.RespondJSON(w, http.StatusOK, pvzArr)
}
