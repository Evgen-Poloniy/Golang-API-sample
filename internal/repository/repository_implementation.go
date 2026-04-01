package repository

type RepositoryImplementation struct {
	connection any
}

func NewRepositoryImplementation(connection any) *RepositoryImplementation {
	return &RepositoryImplementation{connection: connection}
}
