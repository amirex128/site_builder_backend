package middlewares

import (
	"site_builder_backend/pkg/logger"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func buildRequestMessage(c *gin.Context) string {
	var result strings.Builder

	result.WriteString(c.ClientIP())
	result.WriteString(" - ")
	result.WriteString(c.Request.Method)
	result.WriteString(" ")
	result.WriteString(c.Request.URL.String())
	result.WriteString(" - ")
	result.WriteString(strconv.Itoa(c.Writer.Status()))
	result.WriteString(" ")
	result.WriteString(strconv.Itoa(c.Writer.Size()))

	return result.String()
}

// LoggerMiddleware returns a Gin middleware for logging HTTP requests
func LoggerMiddleware(l logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Process request
		c.Next()

		// Log after request
		duration := time.Since(start)
		l.Info("%s - %s", buildRequestMessage(c), duration)
	}
}
