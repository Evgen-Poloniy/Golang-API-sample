package database

import (
	"api_sample/internal/config"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func NewPostgreSQL(config *config.DatabaseConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.Username, config.Password, config.DBName, config.SSLMode,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open db: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping db: %w", err)
	}

	// queries := []string{}

	// for _, query := range queries {
	// 	if _, err := db.Exec(query); err != nil {
	// 		return nil, fmt.Errorf("failed to execute query: %w", err)
	// 	}
	// }

	// query := string

	// if _, err := db.Exec(query); err != nil {
	// 	return nil, fmt.Errorf("failed to execute query: %w", err)
	// }

	return db, nil
}
