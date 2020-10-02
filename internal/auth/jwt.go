package auth

import (
	"errors"
	"github.com/code7unner/leadersofdigital2020-backend/internal/controller/model"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
)

type AuthToken struct {
	TokenType string `json:"token_type"`
	Token     string `json:"access_token"`
	ExpiresIn int64  `json:"expires_in"`
}

type TokenClaim struct {
	*jwt.StandardClaims
	model.RegisterUser
}

func VerifyRequest(r *http.Request, findTokenFns ...func(r *http.Request) string) (string, error) {
	var tokenStr string
	for _, fn := range findTokenFns {
		tokenStr = fn(r)
		if tokenStr != "" {
			break
		}
	}
	if tokenStr == "" {
		return "", errors.New("jwtauth: no token found")
	}

	return tokenStr, nil
}

// TokenFromCookie tries to retreive the token string from a cookie named "jwt".
func TokenFromCookie(r *http.Request) string {
	cookie, err := r.Cookie("jwt")
	if err != nil {
		return ""
	}
	return cookie.Value
}

// "Authorization" reqeust header: "Authorization: BEARER T".
func TokenFromHeader(r *http.Request) string {
	// Get token from authorization header.
	bearer := r.Header.Get("Authorization")
	if len(bearer) > 7 && strings.ToUpper(bearer[0:6]) == "BEARER" {
		return bearer[7:]
	}
	return ""
}

// TokenFromQuery tries to retreive the token string from the "jwt" URI query parameter.
func TokenFromQuery(r *http.Request) string {
	// Get token from query param named "jwt".
	return r.URL.Query().Get("jwt")
}
