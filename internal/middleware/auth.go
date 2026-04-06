package middleware

import (
	errs "api_sample/pkg/errors"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func APIKeyAuth(apiKeyHash []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.Error(errs.ErrMissingAuthHeader)
			c.Abort()
			return
		}

		const prefix = "API-KEY "
		if !strings.HasPrefix(authHeader, prefix) {
			c.Error(errs.ErrWrongAuthHeader)
			c.Abort()
			return
		}

		apiKey := strings.TrimPrefix(authHeader, prefix)

		if bcrypt.CompareHashAndPassword(apiKeyHash, []byte(apiKey)) != nil {
			c.Error(errs.ErrInvalidAPIKey)
			c.Abort()
			return
		}

		c.Next()
	}
}
