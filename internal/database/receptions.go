package database

import (
	"context"
	"github.com/Egorrrad/avitotechBackend2025/internal/dto"
)

func (r *Repository) AddReception(ctx context.Context, user UserModel) (dto.Reception, error) {
	return dto.Reception{}, nil
}

func (r *Repository) closeReception(ctx context.Context, user UserModel) error {
	return nil
}
