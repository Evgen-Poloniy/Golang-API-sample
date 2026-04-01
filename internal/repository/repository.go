package repository

type RepositoryInterface interface {
}

type RepositoryStruct struct {
	RepositoryInterface
}

func NewRepository(connection any) *RepositoryStruct {
	return &RepositoryStruct{
		RepositoryInterface: NewRepositoryImplementation(connection),
	}
}
