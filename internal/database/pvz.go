package database

import (
	"context"
)

func (r *Repository) AddPVZ(ctx context.Context, pvz PvzModel) error {
	stmt := `insert into pvz (id, city, registration_date) values ($1, $2, $3)`
	_, err := r.db.DB.Exec(ctx, stmt, pvz.Id, pvz.City, pvz.RegistrationDate)
	if err != nil {
		return err
	}
	return nil
}
