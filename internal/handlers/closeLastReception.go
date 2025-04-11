package handlers

import (
	"avitotechBackend2025/internal/dto"
	h "avitotechBackend2025/internal/pkg/http"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"net/http"
	"time"
)

// PvzIdCloseLastReception Закрытие последней открытой приемки товаров в рамках ПВЗ
func PvzIDCloseLastReception(w http.ResponseWriter, r *http.Request) {
	pvzIDParam := chi.URLParam(r, "pvzID")

	pvzID, err := uuid.Parse(pvzIDParam)

	if err != nil {
		h.SendError(w, http.StatusBadRequest, "Invalid PVZ ID")
	}

	// логuика закрытия пвз

	response := dto.Reception{
		DateTime: time.Time{},
		Id:       nil,
		PvzId:    pvzID,
		Status:   "closed",
	}

	h.RespondJSON(w, http.StatusOK, response)
}
