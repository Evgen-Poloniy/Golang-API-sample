package repository

import (
	"api_sample/pkg/logger"
	"database/sql"
)

type Repository interface {
}

type RepositoryStorage struct {
	Repository
}

func NewRepository(db *sql.DB, logger logger.Logger) *RepositoryStorage {
	return &RepositoryStorage{
		Repository: NewPostgresRepository(db, logger),
	}
}

// func NewRepository(client *mongo.Client) *RepositoryStorage {
// 	return &RepositoryStorage{
// 		Repository: NewMongoRepository(client),
// 	}
// }

// func NewRepository(client *redis.Client) *RepositoryStorage {
// 	return &RepositoryStorage{
// 		Repository: NewRedisRepository(client),
// 	}
// }
