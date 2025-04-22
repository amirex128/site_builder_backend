package middlewares

import (
	"fmt"
	"runtime/debug"
	"site_builder_backend/pkg/logger"
	"strings"

	"github.com/gin-gonic/gin"
)

// Captured from the logger import that might be available
type Logger interface {
	Error(msg string)
}

func buildPanicMessage(c *gin.Context, err interface{}) string {
	var result strings.Builder

	result.WriteString(c.ClientIP())
	result.WriteString(" - ")
	result.WriteString(c.Request.Method)
	result.WriteString(" ")
	result.WriteString(c.Request.URL.String())
	result.WriteString(" PANIC DETECTED: ")
	result.WriteString(fmt.Sprintf("%v\n%s\n", err, debug.Stack()))

	return result.String()
}

// RecoveryMiddleware returns a middleware that recovers from panics
func RecoveryMiddleware(l logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Log the error
				l.Error(buildPanicMessage(c, err))

				// Respond with Internal Server Error
				c.AbortWithStatus(500)
			}
		}()

		c.Next()
	}
}
