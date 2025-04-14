package repository

import (
	"context"
	"github.com/Egorrrad/avitotechBackend2025/internal/database"
	"github.com/Egorrrad/avitotechBackend2025/internal/dto"
)

func (r *DataManager) AddPVZ(ctx context.Context, pvzReq dto.PostPvzJSONRequestBody) (*dto.PVZ, error) {
	pvz := dto.PVZ{
		City:             pvzReq.City,
		Id:               pvzReq.Id,
		RegistrationDate: pvzReq.RegistrationDate,
	}
	pvzModel := database.PvzModel{
		PVZ: pvz,
	}

	err := r.db.AddPVZ(ctx, pvzModel)
	if err != nil {
		return nil, err
	}

	return &pvz, nil
}
