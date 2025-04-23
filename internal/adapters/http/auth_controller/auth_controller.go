package auth_controller

import (
	"net/http"
	"time"

	authImpl "site_builder_backend/internal/infrastructures/impl/auth"
	"site_builder_backend/internal/interfaces/auth_inter"
	"site_builder_backend/pkg/logger"

	"github.com/gin-gonic/gin"
)

// AuthController handles authentication-related endpoints
type AuthController struct {
	jwtService auth_inter.JWTService
	l          *logger.ZapLogger
}

// LoginRequest holds the user credentials for login
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// RegisterRequest holds the user details for registration
type RegisterRequest struct {
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=6"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
}

// NewAuthController creates a new auth_inter controller
func NewAuthController(jwtService auth_inter.JWTService, l *logger.ZapLogger) *AuthController {
	return &AuthController{
		jwtService: jwtService,
		l:          l,
	}
}

// Login authenticates a user and returns JWT tokens
func (a *AuthController) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Here you would validate user credentials against your database
	// For example: user, err := a.userService.ValidateCredentials(c, req.Email, req.Password)
	// If invalid, return unauthorized

	// For this example, we'll assume the user is valid and create tokens
	// In a real app, you would use the actual user ID and role from the database
	claimsBuilder := authImpl.NewClaimsBuilder().
		WithSubject("user123"). // Replace with actual user ID
		WithCustomClaim("email", req.Email).
		WithRole("user") // Set the user's role

	tokenResponse, err := a.jwtService.Generate(c.Request.Context(), claimsBuilder)
	if err != nil {
		a.l.Error("Failed to generate token", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Authentication failed"})
		return
	}

	c.JSON(http.StatusOK, tokenResponse)
}

// RefreshToken generates new tokens from a refresh token
func (a *AuthController) RefreshToken(c *gin.Context) {
	var req struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokenResponse, err := a.jwtService.RefreshToken(c.Request.Context(), req.RefreshToken)
	if err != nil {
		a.l.Error("Failed to refresh token", "error", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}

	c.JSON(http.StatusOK, tokenResponse)
}

// GetProtectedResource is a sample protected endpoint
func (a *AuthController) GetProtectedResource(c *gin.Context) {
	// The auth_inter middleware already validated the token and set claims in context
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Claims not found in context"})
		return
	}

	userID, _ := c.Get("user_id")
	userRole, _ := c.Get("user_role")

	c.JSON(http.StatusOK, gin.H{
		"message":   "You have access to this protected resource",
		"user_id":   userID,
		"user_role": userRole,
		"timestamp": time.Now(),
		"claims":    claims,
	})
}
