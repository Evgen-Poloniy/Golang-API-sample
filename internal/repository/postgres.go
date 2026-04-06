package repository

import (
	"api_sample/pkg/logger"
	"database/sql"
)

type PostgresRepository struct {
	db     *sql.DB
	logger logger.Logger
}

func NewPostgresRepository(db *sql.DB, logger logger.Logger) *PostgresRepository {
	return &PostgresRepository{
		db:     db,
		logger: logger,
	}
}
