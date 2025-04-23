package auth_inter

import (
	"context"
	"time"
)

// Claims represents the JWT claims
type Claims map[string]interface{}

// TokenType defines the type of token
type TokenType string

const (
	// AccessToken is used for regular API access
	AccessToken TokenType = "access"
	// RefreshToken is used to obtain new access tokens
	RefreshToken TokenType = "refresh"
)

// TokenResponse holds the generated tokens
type TokenResponse struct {
	AccessToken           string        `json:"access_token"`
	RefreshToken          string        `json:"refresh_token"`
	AccessTokenExpiresIn  time.Duration `json:"access_token_expires_in"`
	RefreshTokenExpiresIn time.Duration `json:"refresh_token_expires_in"`
}

// JWTService defines the interface for JWT token operations
type JWTService interface {
	// Generate creates access and refresh tokens using the claims builder
	Generate(ctx context.Context, claimsBuilder ClaimsBuilder) (*TokenResponse, error)

	// ValidateToken validates the given token and returns its claims
	ValidateToken(ctx context.Context, token string, tokenType TokenType) (Claims, error)

	// RefreshToken generates new tokens using a valid refresh token
	RefreshToken(ctx context.Context, refreshToken string) (*TokenResponse, error)

	// GetClaim extracts a specific claim from a token
	GetClaim(ctx context.Context, token string, claimKey string) (interface{}, error)
}

// ClaimsBuilder is an interface for building JWT claims using the builder pattern
type ClaimsBuilder interface {
	// Build creates the final claims
	Build() Claims

	// WithSubject sets the subject claim (usually user ID)
	WithSubject(subject string) ClaimsBuilder

	// WithIssuer sets the issuer claim
	WithIssuer(issuer string) ClaimsBuilder

	// WithAudience sets the audience claim
	WithAudience(audience string) ClaimsBuilder

	// WithExpiresAt sets the expiration time
	WithExpiresAt(expiresAt time.Time) ClaimsBuilder

	// WithIssuedAt sets the issued at time
	WithIssuedAt(issuedAt time.Time) ClaimsBuilder

	// WithNotBefore sets the not before time
	WithNotBefore(notBefore time.Time) ClaimsBuilder

	// WithID sets the JWT ID
	WithID(id string) ClaimsBuilder

	// WithCustomClaim adds a custom claim
	WithCustomClaim(key string, value interface{}) ClaimsBuilder
}
