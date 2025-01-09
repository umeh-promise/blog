package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID int64) (string, error) {

	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(tokenExp).Unix(),
		"iat": time.Now().Unix(),
		"nbf": time.Now().Unix(),
		"iss": tokenIssuer,
		"aud": tokenIssuer,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(authSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpcted signing method %v", t.Header["alg"])
		}

		return []byte(authSecret), nil
	},
		jwt.WithExpirationRequired(),
		jwt.WithAudience(tokenIssuer),
		jwt.WithIssuer(tokenIssuer),
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Name}),
	)
}
