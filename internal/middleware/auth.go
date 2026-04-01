package middleware

import (
	"errors"
	"project/internal/dto"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func APIKeyAuth(apiKeyHash []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			err := errors.New("Authorization header is required")
			c.Error(err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse{
				Error: dto.Error{
					Code:    "UNAUTHORIZED",
					Message: err.Error(),
				},
			})
			return
		}

		const prefix = "API-KEY "
		if !strings.HasPrefix(authHeader, prefix) {
			err := errors.New("Authorization header must start with 'API-KEY '")
			c.Error(err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse{
				Error: dto.Error{
					Code:    "UNAUTHORIZED",
					Message: err.Error(),
				},
			})
			return
		}

		apiKey := strings.TrimPrefix(authHeader, prefix)

		if bcrypt.CompareHashAndPassword(apiKeyHash, []byte(apiKey)) != nil {
			err := errors.New("invalid API key")
			c.Error(err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse{
				Error: dto.Error{
					Code:    "UNAUTHORIZED",
					Message: err.Error(),
				},
			})
			return
		}

		c.Next()
	}
}
