package handlers

import (
	pkg "github.com/Egorrrad/avitotechBackend2025/internal/pkg/http"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"net/http"
)

// PvzIDDeleteLastProduct Удаление последнего добавленного товара из текущей приемки (LIFO, только для сотрудников ПВЗ)
func (h *Handler) PvzIDDeleteLastProduct(w http.ResponseWriter, r *http.Request) {
	pvzIDParam := chi.URLParam(r, "pvzId")

	pvzID, err := uuid.Parse(pvzIDParam)

	// нужно проверить, что  или не нужносо

	if err != nil {
		pkg.SendError(w, http.StatusBadRequest, "Invalid PVZ ID")
	}

	err := h.repository.DeleteLastProduct(r.Context(), pvzID)
	if err != nil {
		return
	}

	pkg.RespondJSON(w, http.StatusOK, pvzID)
}
