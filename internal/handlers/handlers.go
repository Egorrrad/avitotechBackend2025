package handlers

import (
	"github.com/Egorrrad/avitotechBackend2025/internal/repository"
	"log/slog"
)

type Handler struct {
	repository repository.DataManager
	logger     *slog.Logger
}

func New(r repository.DataManager, l *slog.Logger) Handler {
	return Handler{
		repository: r,
		logger:     l,
	}
}
