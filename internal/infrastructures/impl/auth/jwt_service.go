package auth

import (
	"context"
	"errors"
	"fmt"
	"site_builder_backend/configs"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	authInterface "site_builder_backend/internal/interfaces/auth_inter"
)

// JWTConfig holds the configuration for the JWT service
type JWTConfig struct {
	AccessTokenSecret      string
	RefreshTokenSecret     string
	AccessTokenExpiration  time.Duration
	RefreshTokenExpiration time.Duration
	Issuer                 string
}

// jwtService implements the auth_inter.JWTService interface
type jwtService struct {
	config JWTConfig
}

// NewJWTService creates a new instance of the JWT service
func NewJWTService(cfg *configs.Config) authInterface.JWTService {
	return &jwtService{
		config: JWTConfig{
			AccessTokenSecret:      cfg.JWT.AccessTokenSecret,
			RefreshTokenSecret:     cfg.JWT.RefreshTokenSecret,
			AccessTokenExpiration:  cfg.JWT.AccessTokenExpiration,
			RefreshTokenExpiration: cfg.JWT.RefreshTokenExpiration,
			Issuer:                 cfg.JWT.Issuer,
		},
	}
}

// Generate creates access and refresh tokens using the claims builder
func (s *jwtService) Generate(ctx context.Context, claimsBuilder authInterface.ClaimsBuilder) (*authInterface.TokenResponse, error) {
	// Generate access token
	accessClaims := claimsBuilder.
		WithIssuer(s.config.Issuer).
		WithIssuedAt(time.Now()).
		WithExpiresAt(time.Now().Add(s.config.AccessTokenExpiration)).
		WithID(uuid.New().String()).
		WithCustomClaim("type", string(authInterface.AccessToken)).
		Build()

	accessToken, err := s.generateToken(accessClaims, s.config.AccessTokenSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to generate access token: %w", err)
	}

	// Generate refresh token with minimal claims
	refreshClaims := NewClaimsBuilder().
		WithSubject(accessClaims["sub"].(string)).
		WithIssuer(s.config.Issuer).
		WithIssuedAt(time.Now()).
		WithExpiresAt(time.Now().Add(s.config.RefreshTokenExpiration)).
		WithID(uuid.New().String()).
		WithCustomClaim("type", string(authInterface.RefreshToken)).
		Build()

	refreshToken, err := s.generateToken(refreshClaims, s.config.RefreshTokenSecret)
	if err != nil {
		return nil, fmt.Errorf("failed to generate refresh token: %w", err)
	}

	return &authInterface.TokenResponse{
		AccessToken:           accessToken,
		RefreshToken:          refreshToken,
		AccessTokenExpiresIn:  s.config.AccessTokenExpiration,
		RefreshTokenExpiresIn: s.config.RefreshTokenExpiration,
	}, nil
}

// ValidateToken validates the given token and returns its claims
func (s *jwtService) ValidateToken(ctx context.Context, tokenString string, tokenType authInterface.TokenType) (authInterface.Claims, error) {
	var secret string
	if tokenType == authInterface.AccessToken {
		secret = s.config.AccessTokenSecret
	} else {
		secret = s.config.RefreshTokenSecret
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing algorithm
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid claims format")
	}

	// Verify token type
	if claimType, exists := claims["type"]; exists {
		if claimType != string(tokenType) {
			return nil, errors.New("token type mismatch")
		}
	} else {
		return nil, errors.New("token type not found in claims")
	}

	return authInterface.Claims(claims), nil
}

// RefreshToken generates new tokens using a valid refresh token
func (s *jwtService) RefreshToken(ctx context.Context, refreshToken string) (*authInterface.TokenResponse, error) {
	// Validate the refresh token
	claims, err := s.ValidateToken(ctx, refreshToken, authInterface.RefreshToken)
	if err != nil {
		return nil, fmt.Errorf("invalid refresh token: %w", err)
	}

	// Create a new claims builder with the subject from the refresh token
	subject, _ := claims["sub"].(string)
	claimsBuilder := NewClaimsBuilder().WithSubject(subject)

	// Generate new tokens
	return s.Generate(ctx, claimsBuilder)
}

// GetClaim extracts a specific claim from a token
func (s *jwtService) GetClaim(ctx context.Context, token string, claimKey string) (interface{}, error) {
	claims, err := s.ValidateToken(ctx, token, authInterface.AccessToken)
	if err != nil {
		return nil, err
	}

	value, exists := claims[claimKey]
	if !exists {
		return nil, fmt.Errorf("claim %s not found", claimKey)
	}

	return value, nil
}

// generateToken creates a new signed token from claims
func (s *jwtService) generateToken(claims authInterface.Claims, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(claims))
	return token.SignedString([]byte(secret))
}

// Ensure jwtService implements the JWTService interface
var _ authInterface.JWTService = (*jwtService)(nil)
