package main

import (
	"context"
	"github.com/Egorrrad/avitotechBackend2025/configs"
	"log/slog"
)

func main() {
	ctx := context.Background()
	cfg, err := configs.Load()
	if err != nil {
		slog.Error("FATAL", "error", err)
	}

	// store := store.NewMemoryMoviesStore()
	server, err := NewServer(cfg)
	if err != nil {
		slog.Error("Error when server init", "error", err)
	}
	server.Start(ctx)
}
