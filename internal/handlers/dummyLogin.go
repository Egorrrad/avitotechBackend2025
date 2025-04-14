package handlers

import (
	"encoding/json"
	"github.com/Egorrrad/avitotechBackend2025/internal/dto"
	pkg "github.com/Egorrrad/avitotechBackend2025/internal/pkg/http"
	token "github.com/Egorrrad/avitotechBackend2025/internal/pkg/jwt"
	"net/http"
)

func (h *Handler) DummyLogin(w http.ResponseWriter, r *http.Request) {
	var req dto.PostDummyLoginJSONRequestBody

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		pkg.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	role := string(req.Role)
	tokenString, err := token.GenerateToken(role)
	if err != nil {
		pkg.SendError(w, http.StatusInternalServerError, pkg.ErrFailedGenerateToken)
		return
	}

	response := tokenString

	pkg.RespondJSON(w, http.StatusOK, response)
}
