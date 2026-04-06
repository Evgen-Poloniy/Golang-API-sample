package service

import (
	"api_sample/internal/repository"
	"api_sample/pkg/logger"
)

type Service interface {
}

type ServiceUsecase struct {
	Service
}

func NewService(repository repository.Repository, logger logger.Logger) *ServiceUsecase {
	return &ServiceUsecase{
		Service: NewTODOService(repository, logger),
	}
}
