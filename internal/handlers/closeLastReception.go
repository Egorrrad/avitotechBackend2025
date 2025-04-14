package handlers

import (
	pkg "github.com/Egorrrad/avitotechBackend2025/internal/pkg/http"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"net/http"
	_ "time"
)

// PvzIdCloseLastReception Закрытие последней открытой приемки товаров в рамках ПВЗ
func (h *Handler) PvzIDCloseLastReception(w http.ResponseWriter, r *http.Request) {
	pvzIDParam := chi.URLParam(r, "pvzID")

	pvzID, err := uuid.Parse(pvzIDParam)

	if err != nil {

		pkg.SendError(w, http.StatusBadRequest, "Invalid PVZ ID")
	}

	// логuика закрытия пвз

	reception, err := h.Repository.CloseReception(r.Context(), pvzID)

	if err != nil {
		pkg.SendError(w, http.StatusInternalServerError, err.Error())
	}
	pkg.RespondJSON(w, http.StatusOK, reception)
}
