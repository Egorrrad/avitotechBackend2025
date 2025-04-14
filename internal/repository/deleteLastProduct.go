package repository

import (
	"context"
	"errors"
	"github.com/oapi-codegen/runtime/types"
)

func (r *DataManager) DeleteLastProduct(ctx context.Context, pvzID types.UUID) error {
	// надо найти сначала незакрытую приемку
	// если она закрыта6 то ошибка

	unclsReception, err := r.db.getLastUnclsReception(ctx, pvzID)
	if err != nil {
		return err
	}
	if unclsReception == nil {
		return errors.New("no uncls reception found")
	}

	err = r.db.DeleteLastProduct(ctx, *unclsReception.Id)

	if err != nil {
		return err
	}

	return nil
}
