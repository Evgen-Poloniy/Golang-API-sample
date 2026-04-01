package middleware

import (
	"fmt"
	"os"
	"project/internal/config"
	"project/internal/logutil"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func InitLoggingMiddleware(cfg *config.LoggerConfig) gin.HandlerFunc {
	logger := logrus.New()
	logger.SetOutput(os.Stdout)
	logutil.SetLevel(logger, cfg)
	logutil.SetFormatter(logger, cfg)

	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()

		entry := logger.WithFields(logrus.Fields{
			"status":  status,
			"method":  c.Request.Method,
			"path":    c.Request.URL.Path,
			"ip":      c.ClientIP(),
			"latency": fmt.Sprintf("%v", latency),
		})

		if len(c.Errors) > 0 {
			err := c.Errors.Last()
			code := ""
			message := strings.Join(c.Errors.Errors(), "; ")

			if meta, ok := err.Meta.(map[string]interface{}); ok {
				if codeVal, exists := meta["code"]; exists {
					if str, ok := codeVal.(string); ok {
						code = str
					}
				}
				if messageVal, exists := meta["message"]; exists {
					if str, ok := messageVal.(string); ok && str != "" {
						message = str
					}
				}
			}

			entry = entry.WithFields(logrus.Fields{
				"code": code,
			})

			if status >= 500 {
				entry.Error(message)
			} else {
				entry.Warn(message)
			}

			return
		}

		entry.Info("request completed")
	}
}
