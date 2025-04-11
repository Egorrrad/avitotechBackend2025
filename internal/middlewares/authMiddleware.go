package middlewares

import (
	"avitotechBackend2025/internal/dto"
	helper "avitotechBackend2025/internal/pkg/http"
	token "avitotechBackend2025/internal/pkg/jwt"
	"context"
	"net/http"
	"strings"
)

func AuthMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// проверяем токен пользователя и наверно добавим пользователя потом в контекст или только его роль

		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			helper.SendError(w, http.StatusForbidden, helper.ErrHeaderIsRequired)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			helper.SendError(w, http.StatusForbidden, helper.ErrInvalidTokenFormat)
			return
		}

		tokenString := parts[1]

		claims, err := token.ParseToken(tokenString)
		if err != nil {
			helper.SendError(w, http.StatusForbidden, helper.ErrInvalidToken)
		}

		ctx := context.WithValue(r.Context(), dto.UserRoleKey, claims.Role)
		handler.ServeHTTP(w, r.WithContext(ctx))
	})

}
