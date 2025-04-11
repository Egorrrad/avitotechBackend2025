package handlers

import (
	"avitotechBackend2025/internal/dto"
	h "avitotechBackend2025/internal/pkg/http"
	token "avitotechBackend2025/internal/pkg/jwt"
	"encoding/json"
	"net/http"
)

func DummyLogin(w http.ResponseWriter, r *http.Request) {
	var req dto.PostDummyLoginJSONRequestBody

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	role := string(req.Role)
	tokenString, err := token.GenerateToken(role)
	if err != nil {
		h.SendError(w, http.StatusInternalServerError, h.ErrFailedGenerateToken)
		return
	}

	response := tokenString

	h.RespondJSON(w, http.StatusOK, response)
}
