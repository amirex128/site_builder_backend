package middlewares

import (
	"net/http"
	"strings"

	"site_builder_backend/internal/interfaces/auth_inter"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware provides JWT authentication middleware for Gin
type AuthMiddleware struct {
	jwtService auth_inter.JWTService
}

// NewAuthMiddleware creates a new auth_inter middleware
func NewAuthMiddleware(jwtService auth_inter.JWTService) *AuthMiddleware {
	return &AuthMiddleware{
		jwtService: jwtService,
	}
}

// Authenticate middleware verifies JWT token and sets user claims in context
func (m *AuthMiddleware) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authorization header is required"})
			return
		}

		// Check if the header has the format "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header format"})
			return
		}

		tokenString := parts[1]
		claims, err := m.jwtService.ValidateToken(c.Request.Context(), tokenString, auth_inter.AccessToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token", "details": err.Error()})
			return
		}

		// Set user claims in context for later use
		c.Set("claims", claims)
		c.Set("user_id", claims["sub"])
		if role, exists := claims["role"]; exists {
			c.Set("user_role", role)
		}

		c.Next()
	}
}
