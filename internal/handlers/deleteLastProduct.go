package handlers

import (
	h "avitotechBackend2025/internal/pkg/http"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"net/http"
)

// PvzIDDeleteLastProduct Удаление последнего добавленного товара из текущей приемки (LIFO, только для сотрудников ПВЗ)
func PvzIDDeleteLastProduct(w http.ResponseWriter, r *http.Request) {
	pvzIDParam := chi.URLParam(r, "pvzId")

	pvzID, err := uuid.Parse(pvzIDParam)

	// нужно проверить, что  или не нужносо

	if err != nil {
		h.SendError(w, http.StatusBadRequest, "Invalid PVZ ID")
	}

	h.RespondJSON(w, http.StatusOK, pvzID)
}
