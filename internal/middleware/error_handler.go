package middleware

import (
	"api_sample/internal/dto"
	errs "api_sample/pkg/errors"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) == 0 {
			return
		}

		var statusCode int
		var code string
		err := c.Errors.Last().Err

		// Error mapping
		if appError, ok := errors.AsType[*errs.AppError](err); ok {
			statusCode = appError.StatusCode
			code = appError.Code
		} else if errors.Is(err, errs.ErrMissingAuthHeader) {
			statusCode = http.StatusUnauthorized
			code = "MISSING_AUTH_HEADER"
		} else if errors.Is(err, errs.ErrWrongAuthHeader) {
			statusCode = http.StatusUnauthorized
			code = "WRONG_AUTH_HEADER"
		} else if errors.Is(err, errs.ErrInvalidAPIKey) {
			statusCode = http.StatusUnauthorized
			code = "INVALID_API_KEY"
		} else {
			statusCode = http.StatusInternalServerError
			code = "UNKNOWN_ERROR"
		}

		c.Set("status_code", statusCode)
		c.Set("code", code)

		c.JSON(statusCode, dto.ErrorResponse{
			Error: dto.Error{
				Code:    code,
				Message: err.Error(),
			},
		})
	}
}
