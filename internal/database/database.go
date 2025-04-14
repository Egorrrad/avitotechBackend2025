package database

import (
	"context"
	"fmt"
	"github.com/Egorrrad/avitotechBackend2025/configs"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

type Database struct {
	DB *pgxpool.Pool
}

func NewDatabase(db *pgxpool.Pool) Database {
	return Database{
		DB: db,
	}
}

type Repository struct {
	db Database
}

func NewRepository(db Database) *Repository {
	return &Repository{
		db: db,
	}
}

func ConnectDB(cfg *configs.PostgresConf) *Database {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Password,
		cfg.Name,
	)
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatal("Unable to parse DB config: %v\n", err)
	}

	dbpool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	db := &Database{DB: dbpool}

	return db
}
