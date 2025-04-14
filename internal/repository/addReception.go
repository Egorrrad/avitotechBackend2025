package repository

import (
	"context"
	"github.com/Egorrrad/avitotechBackend2025/internal/dto"
)

func (r *DataManager) AddReception(ctx context.Context, reception dto.PostReceptionsJSONRequestBody) (*dto.Reception, error) {
	receptionModel, err := r.db.AddReception(ctx, reception)
	if err != nil {
		return nil, err
	}
	return &receptionModel, nil
}
