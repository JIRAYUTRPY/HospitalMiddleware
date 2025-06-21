package pkg

import (
	"errors"
	"time"

	"github.com/agnos/hospital-middleware/config"
	"github.com/golang-jwt/jwt"
)

var (
	ErrExpiredToken = errors.New("expired token")
)

type Claims struct {
	AccountID int `json:"account_id"`
	jwt.StandardClaims
}

func GenerateTokens(accountID int, jwtConfig config.JWTConfig) (accessToken string, err error) {
	atClaims := &Claims{
		AccountID: accountID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(3 * time.Hour).Unix(),
		},
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	accessToken, err = at.SignedString([]byte(jwtConfig.AccessSecret))
	if err != nil {
		return
	}
	return
}

func ValidateToken(tokenStr string, jwtConfig config.JWTConfig) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtConfig.AccessSecret), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, err
	}
	if time.Now().Unix()-claims.ExpiresAt > 24*60*60 {
		return nil, ErrExpiredToken
	}
	return claims, nil
}
