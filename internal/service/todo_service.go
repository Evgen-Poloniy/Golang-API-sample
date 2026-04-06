package service

import (
	"api_sample/internal/repository"
	"api_sample/pkg/logger"
)

type TODOService struct {
	repository repository.Repository
	logger     logger.Logger
}

func NewTODOService(repository repository.Repository, logger logger.Logger) *TODOService {
	return &TODOService{
		repository: repository,
		logger:     logger,
	}
}
