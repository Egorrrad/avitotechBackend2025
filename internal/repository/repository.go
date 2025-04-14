package repository

import (
	"context"
	"github.com/Egorrrad/avitotechBackend2025/internal/database"
	"github.com/Egorrrad/avitotechBackend2025/internal/dto"
	"github.com/google/uuid"
	"github.com/oapi-codegen/runtime/types"
)

type DataManager struct {
	db DBManager
}

func New(db DBManager) *DataManager {
	return &DataManager{
		db: db,
	}
}

type DBManager interface {
	AddPVZ(ctx context.Context, pvz database.PvzModel) error
	AddReception(ctx context.Context, reception dto.PostReceptionsJSONRequestBody) (dto.Reception, error)
	AddProduct(ctx context.Context, unclsReception dto.Reception, product dto.PostProductsJSONRequestBody) (dto.Product, error)
	getLastUnclsReception(ctx context.Context, pvzID types.UUID) (*dto.Reception, error)
	DeleteLastProduct(ctx context.Context, reception types.UUID) error
	GetLastReception(ctx any, pvzID uuid.UUID) (*dto.Reception, error)
	CloseReception(ctx any, pvzID uuid.UUID) (*dto.Reception, error)
	GetAllPvz(ctx context.Context, reqPVZ dto.GetPvzParams) ([]dto.PVZ, error)
}
