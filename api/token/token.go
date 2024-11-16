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

func GenerateAccessToken(tokenReq model.Token) (string, error) {
	claims := TokenClaims{
		ID:       tokenReq.ID,
		Username: tokenReq.Username,
		Role:     tokenReq.Role,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Minute * 20).Unix(), // Access token muddati qisqaroq
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.LoadConfig().ACCESS_SECRET_KEY))
}

func GenerateRefreshToken(tokenReq model.Token) (string, error) {
	claims := TokenClaims{
		ID:       tokenReq.ID,
		Username: tokenReq.Username,
		Role:     tokenReq.Role,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(), // Refresh token muddati uzoqroq
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.LoadConfig().REFRESH_SECRET_KEY))
}

func ExtractAccessClaims(tokenStr string) (*TokenClaims, error) {
	return extractClaims(tokenStr, config.LoadConfig().ACCESS_SECRET_KEY)
}

func ExtractRefreshClaims(tokenStr string) (*TokenClaims, error) {
	return extractClaims(tokenStr, config.LoadConfig().REFRESH_SECRET_KEY)
}

func extractClaims(tokenStr string, secret string) (*TokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &TokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
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
