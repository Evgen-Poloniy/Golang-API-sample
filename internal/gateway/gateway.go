package gateway

import "net/http"

type Gateway interface {
}

type HTTPGateway struct {
	Gateway
}

func NewHTTPGateway(client *http.Client) *HTTPGateway {
	return &HTTPGateway{
		Gateway: NewExternalService(client),
	}
}
