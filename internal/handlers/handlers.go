package handlers

import (
	"github.com/Egorrrad/avitotechBackend2025/internal/logger"
	"github.com/Egorrrad/avitotechBackend2025/internal/repository"
)

type Handler struct {
	Repository repository.DataManager
	Logger     *logger.Logger
}

func New(r repository.DataManager, l *logger.Logger) Handler {
	return Handler{
		Repository: r,
		Logger:     l,
	}
}
