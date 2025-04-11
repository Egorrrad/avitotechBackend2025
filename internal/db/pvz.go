package db

import "context"

func CreatePVZ(ctx context.Context, pvz PvzModel) error {
	stmt := `insert into pvzs (id, created_at, updated_at) values ($1, $2, $3)`
}

func GetPVZs(ctx context.Context) {
	stmt := `select id, created_at, updated_at from pvzs order by created_at desc`

}
