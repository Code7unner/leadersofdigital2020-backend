package auth

import (
	"encoding/json"
	"github.com/code7unner/leadersofdigital2020-backend/internal/controller/model"
	"github.com/code7unner/leadersofdigital2020-backend/utils"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
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

func Register(secret string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var registerUser model.RegisterUser

		if err := json.NewDecoder(r.Body).Decode(&registerUser); err != nil {
			utils.ErrorHandler(w, err, http.StatusBadRequest)
			return
		}

		expiresAt := time.Now().Add(time.Hour * 100).Unix()
		token := jwt.New(jwt.SigningMethodHS256)

		token.Claims = &TokenClaim{
			StandardClaims: &jwt.StandardClaims{
				ExpiresAt: expiresAt,
			},
			RegisterUser: registerUser,
		}

		tokenString, err := token.SignedString([]byte(secret))
		if err != nil {
			utils.ErrorHandler(w, err, http.StatusBadRequest)
		}

		data := AuthToken{
			Token:     tokenString,
			TokenType: "Bearer",
			ExpiresIn: expiresAt,
		}
		utils.SuccessHandler(w, http.StatusOK, data)
	}
}
