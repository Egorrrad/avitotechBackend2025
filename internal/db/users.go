package db

import (
	"context"
	"github.com/oapi-codegen/runtime/types"
)

func createUser(ctx context.Context, user UserModel) error {
	stmt := `insert into users (id, email, password_hash, role) values ($1, $2, $3, $4)`
	_, err := db.Exec(ctx, stmt, user.Id, user.Email, user.Password, user.Role)
	if err != nil {
		return err
	}
	return nil
}

func getUserById(ctx context.Context, userId int64) (*UserModel, error) {
	user := &UserModel{}
	stmt := `select (id, email, password_hash, role) from users where id = $1`
	err := db.QueryRow(ctx, stmt, userId).Scan(user.Id, user.Email, user.Password, user.Role)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func getUserByEmail(ctx context.Context, email types.Email) (*UserModel, error) {
	user := &UserModel{}
	stmt := `select id, email, password_hash, role from users where email = $1`
	err := db.QueryRow(ctx, stmt, email).Scan(user.Id, user.Email, user.Password, user.Role)
	if err != nil {
		return nil, err
	}
	return user, nil
}
