package repository

import (
	"errors"
	"github.com/Egorrrad/avitotechBackend2025/internal/dto"
	"github.com/google/uuid"
)

func (r *DataManager) CloseReception(ctx any, pvzID uuid.UUID) (*dto.Reception, error) {
	// В случае, если приёмка товаров уже была закрыта (или приёма товаров в данном ПВЗ ещё не было), то следует вернуть ошибку.
	reception, err := r.db.GetLastReception(ctx, pvzID)
	if err != nil {
		return nil, err
	}
	if reception.Status == dto.Close {
		return nil, errors.New(ErrAlreadyClosed)
	}
	reception, err = r.db.CloseReception(ctx, pvzID)

	if err != nil {
		return nil, err
	}

	return reception, nil
}
