package token

import (
	"fmt"
	"tender/config"
	"tender/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type TokenClaims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

func GenerateToken(tokenReq model.Token) (string, error) {
	claims := TokenClaims{
		ID:       tokenReq.ID,
		Username: tokenReq.Username,
		Role:     tokenReq.Role,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Hour * 3).Unix(), // Access token muddati qisqaroq
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.LoadConfig().SECRET_KEY))
}

func ExtractClaims(tokenStr string) (*TokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &TokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.LoadConfig().SECRET_KEY), nil
	})

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("token is invalid or expired")
	}

	claims, ok := token.Claims.(*TokenClaims)
	if !ok {
		return nil, fmt.Errorf("token claims are invalid")
	}

	return claims, nil
}
