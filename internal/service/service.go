package service

import "project/internal/repository"

type ServiceInterface interface {
}

type ServiceStruct struct {
	ServiceInterface
}

func NewService(repository repository.RepositoryInterface) *ServiceStruct {
	return &ServiceStruct{
		ServiceInterface: NewServiceImplementation(repository),
	}
}
