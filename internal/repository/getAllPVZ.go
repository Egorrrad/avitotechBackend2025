package repository

import (
	"context"
	"github.com/Egorrrad/avitotechBackend2025/internal/dto"
)

func (r *DataManager) GetAllPvz(ctx context.Context, req dto.GetPvzParams) (*[]dto.PVZ, error) {
	// При этом добавить фильтр по дате приёмки товаров, т.е. выводить только те
	//ПВЗ и всю информацию по ним, которые в указанный диапазон времени проводили приёмы товаров.
	pvzArray, err := r.db.GetAllPvz(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pvzArray, nil
}
