package main

import (
	"context"
	"fmt"
	"github.com/Egorrrad/avitotechBackend2025/configs"
	"github.com/Egorrrad/avitotechBackend2025/internal/dto"
	"github.com/Egorrrad/avitotechBackend2025/internal/handlers"
	"github.com/Egorrrad/avitotechBackend2025/internal/logger"
	"github.com/Egorrrad/avitotechBackend2025/internal/middlewares"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type Server struct {
	cfg *configs.Configuration
	// store  store.Interface
	router  *chi.Mux
	logger  *logger.Logger
	handler handlers.Handler
}

func NewServer(cfg *configs.Configuration, handler handlers.Handler) (*Server, error) {
	l, err := logger.New(&cfg.LoggerConf)

	if err != nil {
		return nil, err
	}

	srv := &Server{
		cfg:     cfg,
		handler: handler,
		logger:  l,
	}

	srv.router = srv.InitRouter()

	return srv, nil
}

func (s *Server) Start(ctx context.Context) {
	server := http.Server{
		Addr:         fmt.Sprintf(":%d", s.cfg.Port),
		Handler:      s.router,
		IdleTimeout:  s.cfg.IdleTimeout,
		ReadTimeout:  s.cfg.ReadTimeout,
		WriteTimeout: s.cfg.WriteTimeout,
	}

	shutdownComplete := handleShutdown(func() {
		if err := server.Shutdown(ctx); err != nil {
			log.Printf("server.Shutdown failed: %v\n", err)
		}
	})

	s.logger.Info("Service started", "adr", server.Addr)
	if err := server.ListenAndServe(); err == http.ErrServerClosed {
		<-shutdownComplete
	} else {
		log.Printf("http.ListenAndServe failed: %v\n", err)
	}

	log.Println("Shutdown gracefully")
}

func handleShutdown(onShutdownSignal func()) <-chan struct{} {
	shutdown := make(chan struct{})

	go func() {
		shutdownSignal := make(chan os.Signal, 1)
		signal.Notify(shutdownSignal, os.Interrupt, syscall.SIGTERM)

		<-shutdownSignal

		onShutdownSignal()
		close(shutdown)
	}()

	return shutdown
}

func (s *Server) InitRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middlewares.LoggerMiddleware)
	r.Use(middleware.Recoverer)

	r.Post("/dummyLogin", s.handler.DummyLogin)
	r.Post("/login", s.handler.Login)
	r.Post("/register", s.handler.Register)

	r.Group(func(r chi.Router) {
		r.Use(middlewares.AuthMiddleware)

		// только для модераторов
		r.Group(
			func(r chi.Router) {
				r.Use(middlewares.RoleCheckMiddleware(dto.UserRoleModerator))
				r.Post("/pvz", s.handler.AddPVZ)
			})

		// только для модераторов и сотрудников ПВЗ
		r.Group(
			func(r chi.Router) {
				r.Use(middlewares.RoleCheckMiddleware(dto.UserRoleModerator, dto.UserRoleEmployee))
				r.Get("/pvz", s.handler.GetPVZ)
			})

		// только для сотрудников ПВЗ
		r.Group(
			func(r chi.Router) {
				r.Use(middlewares.RoleCheckMiddleware(dto.UserRoleEmployee))
				r.Post("/products", s.handler.Products)
				r.Post("/pvz/{pvzId}/delete_last_product", s.handler.PvzIDDeleteLastProduct)
				r.Post("/pvz/{pvzId}/close_last_reception", s.handler.PvzIDCloseLastReception)
				r.Post("/receptions", s.handler.Receptions)
			})
	})

	return r
}
