package routes

import (
	"avitotechBackend2025/internal/dto"
	"avitotechBackend2025/internal/handlers"
	"avitotechBackend2025/internal/middlewares"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func InitRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middlewares.LoggerMiddleware)
	r.Use(middleware.Recoverer)

	r.Post("/dummyLogin", handlers.DummyLogin)
	r.Post("/login", handlers.Login)
	r.Post("/register", handlers.Register)

	r.Group(func(r chi.Router) {
		r.Use(middlewares.AuthMiddleware)

		// только для модераторов
		r.Group(
			func(r chi.Router) {
				r.Use(middlewares.RoleCheckMiddleware(dto.UserRoleModerator))
				r.Post("/pvz", handlers.CreatePVZ)
			})

		r.Get("/pvz", handlers.GetPVZ)

		r.Post("/pvz/{pvzId}/close_last_reception", handlers.PvzIDCloseLastReception)

		// только для сотрудников ПВЗ
		r.Group(
			func(r chi.Router) {
				r.Use(middlewares.RoleCheckMiddleware(dto.UserRoleEmployee))
				r.Post("/products", handlers.Products)
				r.Post("/pvz/{pvzId}/delete_last_product", handlers.PvzIDDeleteLastProduct)
				r.Post("/receptions", handlers.Receptions)
			})
	})

	return r
}
