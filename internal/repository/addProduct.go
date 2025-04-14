package repository

import (
	"context"
	"errors"
	"github.com/Egorrrad/avitotechBackend2025/internal/dto"
)

func (r *DataManager) AddProduct(ctx context.Context, productsReq dto.PostProductsJSONRequestBody) (*dto.Product, error) {
	// находим последнюю незакрытую приемку
	// если нет, то ошибка

	pvzID := productsReq.PvzId
	unclsReception, err := r.db.getLastUnclsReception(ctx, pvzID)
	if err != nil {
		return nil, err
	}
	if unclsReception == nil {
		return nil, errors.New("no uncls reception found")
	}

	// если норм, то добавляем в незакрытую приемку

	product, err := r.db.AddProduct(ctx, *unclsReception, productsReq)
	if err != nil {
		return nil, err
	}

	return &product, nil
}
