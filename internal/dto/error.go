package dto

// Error represents an API error response
// @Description Standard error response structure
type Error struct {
	Code string `json:"code"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error interface{} `json:"error"`
}
