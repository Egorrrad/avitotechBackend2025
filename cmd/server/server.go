package main

import (
	"context"
	"fmt"
	"github.com/Egorrrad/avitotechBackend2025/configs"
	"github.com/Egorrrad/avitotechBackend2025/internal/logger"
	"github.com/Egorrrad/avitotechBackend2025/internal/routes"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type Server struct {
	cfg *configs.Configuration
	// store  store.Interface
	router *chi.Mux
	logger *logger.Logger
}

func NewServer(cfg *configs.Configuration) (*Server, error) {
	srv := &Server{
		cfg: cfg,
		//store:  store,
	}

	srv.router = routes.InitRouter()

	l, err := logger.New(&cfg.LoggerConf)

	srv.logger = l

	if err != nil {
		return nil, err
	}

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
