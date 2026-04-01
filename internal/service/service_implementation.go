package service

import (
	"project/internal/repository"
)

type ServiceImplementation struct {
	repository repository.RepositoryInterface
}

func NewServiceImplementation(repository repository.RepositoryInterface) *ServiceImplementation {
	return &ServiceImplementation{
		repository: repository,
	}
}
