package auth

import (
	"github.com/code7unner/leadersofdigital2020-backend/internal/logging"
	"github.com/code7unner/leadersofdigital2020-backend/utils"
	"github.com/dgrijalva/jwt-go"
	"net/http"
)

func Authenticator(secret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			logger := logging.FromContext(ctx)
			tokenStr, err := VerifyRequest(r, TokenFromQuery, TokenFromHeader, TokenFromCookie)
			if err != nil {
				logger.Errorf("Auth: token not found: %s", err.Error())
				utils.ErrorHandler(w, err, http.StatusUnauthorized)
				return
			}

			claims := &TokenClaim{}
			tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(secret), nil
			})
			if err != nil {
				if err == jwt.ErrSignatureInvalid {
					w.WriteHeader(http.StatusUnauthorized)
					return
				}
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			if !tkn.Valid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}