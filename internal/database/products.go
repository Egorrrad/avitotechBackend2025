package database

import (
	"context"
	"github.com/Egorrrad/avitotechBackend2025/internal/dto"
)

func (r *Repository) AddProduct(ctx context.Context, product dto.PostProductsJSONRequestBody) (dto.Product, error) {
	return dto.Product{}, nil
}

func (r *Repository) deleteProduct(ctx context.Context, user UserModel) error {
	return nil
}
