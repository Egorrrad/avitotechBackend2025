package handlers

import (
	"encoding/json"
	"github.com/Egorrrad/avitotechBackend2025/internal/database"
	"github.com/Egorrrad/avitotechBackend2025/internal/dto"
	pkg "github.com/Egorrrad/avitotechBackend2025/internal/pkg/http"
	token "github.com/Egorrrad/avitotechBackend2025/internal/pkg/jwt"
	"net/http"
)

// Авторизация пользователя
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req dto.PostLoginJSONRequestBody

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		pkg.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	// получаем юзера из бд

	user := database.UserModel{
		User:     dto.User{},
		Password: "",
	}

	role := user.Role

	tokenString, err := token.GenerateToken(string(role))
	if err != nil {
		pkg.SendError(w, http.StatusInternalServerError, pkg.ErrFailedGenerateToken)
		return
	}

	response := tokenString

	pkg.RespondJSON(w, http.StatusOK, response)
}
