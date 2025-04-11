package handlers

import (
	"avitotechBackend2025/internal/db"
	"avitotechBackend2025/internal/dto"
	h "avitotechBackend2025/internal/pkg/http"
	token "avitotechBackend2025/internal/pkg/jwt"
	"encoding/json"
	"net/http"
)

// Авторизация пользователя
func Login(w http.ResponseWriter, r *http.Request) {
	var req dto.PostLoginJSONRequestBody

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	// получаем юзера из бд

	user := db.UserModel{
		User:     dto.User{},
		Password: "",
	}

	role := user.Role

	tokenString, err := token.GenerateToken(string(role))
	if err != nil {
		h.SendError(w, http.StatusInternalServerError, h.ErrFailedGenerateToken)
		return
	}

	response := tokenString

	h.RespondJSON(w, http.StatusOK, response)
}
