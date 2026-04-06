package gateway

import "net/http"

type ExternalService struct {
	client *http.Client
}

func NewExternalService(client *http.Client) *ExternalService {
	return &ExternalService{client: client}
}
