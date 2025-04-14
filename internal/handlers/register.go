package handlers

import (
	"avitotechBackend2025/internal/dto"
	h "avitotechBackend2025/internal/pkg/http"
	"encoding/json"
	"net/http"
)

// Register Регистрация пользователя
func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var req dto.PostRegisterJSONRequestBody

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	// создание пользователя
	user := &dto.User{
		Email: "",
		Id:    nil,
		Role:  "",
	}

	h.RespondJSON(w, http.StatusCreated, user)

}
