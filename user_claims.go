package easclient

import (
	"context"
	"net/http"
	"strings"
)

// UserClaims aggregates all claims that are scoped to an EAS archive user.
type UserClaims struct {
	// UserId - Technischer Identifikator des Nutzers, der im Host-System die Anfrage ausgelöst hat.
	UserId string `json:"user,omitempty"`

	// UserFullname - Optionaler, vollständiger Name des Nutzers, der im Host-System die Anfrage ausgelöst hat
	UserFullname string `json:"user_fullname,omitempty"`

	// Tokens - Securitytokens (Gruppenzugehörigkeit) -getrennt durch Kommata- des Nutzers, der im Host-System den Request
	// ausgelöst hat.
	Tokens []string `json:"tokens,omitempty"`
}

func NewUserClaims(userId string) *UserClaims {
	return &UserClaims{UserId: userId, Tokens: []string{}}
}

// SetOnHeader sets all possible eas.RequestHeader based on this UserClaims
func (claims *UserClaims) SetOnHeader(header http.Header) {
	header.Set(string(HeaderUser), claims.UserId)

	// Set optional fullname
	if claims.UserFullname != "" {
		header.Set(string(HeaderUserFullname), claims.UserFullname)
	}

	// Set optional tokens
	if len(claims.Tokens) > 0 {
		header.Set(string(HeaderTokens), strings.Join(claims.Tokens, ","))
	} else {
		header.Set(string(HeaderTokens), "")
	}
}

const UserClaimsKey = "easclient-user-claims-key"

func (claims *UserClaims) SetOnContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, UserClaimsKey, claims)
}

func UserClaimsFromContext(ctx context.Context) *UserClaims {
	if ctx == nil {
		return nil
	}

	if claims, ok := ctx.Value(UserClaimsKey).(*UserClaims); ok {
		return claims
	}

	return nil
}
