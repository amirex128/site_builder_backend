package auth

import (
	"time"

	authInterface "site_builder_backend/internal/interfaces/auth_inter"
)

// claimsBuilder implements the authInterface.ClaimsBuilder interface
type claimsBuilder struct {
	claims authInterface.Claims
}

// NewClaimsBuilder creates a new claims builder
func NewClaimsBuilder() authInterface.ClaimsBuilder {
	return &claimsBuilder{
		claims: make(authInterface.Claims),
	}
}

// Build creates the final claims
func (b *claimsBuilder) Build() authInterface.Claims {
	return b.claims
}

// WithSubject sets the subject claim (usually user ID)
func (b *claimsBuilder) WithSubject(subject string) authInterface.ClaimsBuilder {
	b.claims["sub"] = subject
	return b
}

// WithIssuer sets the issuer claim
func (b *claimsBuilder) WithIssuer(issuer string) authInterface.ClaimsBuilder {
	b.claims["iss"] = issuer
	return b
}

// WithAudience sets the audience claim
func (b *claimsBuilder) WithAudience(audience string) authInterface.ClaimsBuilder {
	b.claims["aud"] = audience
	return b
}

// WithExpiresAt sets the expiration time
func (b *claimsBuilder) WithExpiresAt(expiresAt time.Time) authInterface.ClaimsBuilder {
	b.claims["exp"] = expiresAt.Unix()
	return b
}

// WithIssuedAt sets the issued at time
func (b *claimsBuilder) WithIssuedAt(issuedAt time.Time) authInterface.ClaimsBuilder {
	b.claims["iat"] = issuedAt.Unix()
	return b
}

// WithNotBefore sets the not before time
func (b *claimsBuilder) WithNotBefore(notBefore time.Time) authInterface.ClaimsBuilder {
	b.claims["nbf"] = notBefore.Unix()
	return b
}

// WithID sets the JWT ID
func (b *claimsBuilder) WithID(id string) authInterface.ClaimsBuilder {
	b.claims["jti"] = id
	return b
}

// WithCustomClaim adds a custom claim
func (b *claimsBuilder) WithCustomClaim(key string, value interface{}) authInterface.ClaimsBuilder {
	b.claims[key] = value
	return b
}

// Ensure claimsBuilder implements the ClaimsBuilder interface
var _ authInterface.ClaimsBuilder = (*claimsBuilder)(nil)
