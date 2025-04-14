package middlewares

import (
	"github.com/Egorrrad/avitotechBackend2025/internal/dto"
	pkg "github.com/Egorrrad/avitotechBackend2025/internal/pkg/http"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

func RoleCheckMiddleware(roles ...dto.UserRole) func(http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			claims := r.Context().Value(dto.UserRoleKey).(jwt.MapClaims) // наверно преобразование типа надо проверить
			allowed := false
			for _, role := range roles {
				if role == claims["role"] {
					allowed = true
					handler.ServeHTTP(w, r)
					break
				}
			}

			if !allowed {
				pkg.SendError(w, http.StatusForbidden, pkg.ErrRoleNotAllowed)
				return
			}
		})
	}
}
