package middleware

import (
	"errors"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(userId string, username string) (string, error) {
	exp := time.Now().Add(time.Hour * 1).Unix()

	claims := payload.ClaimsCustom{
		UserId:   userId,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp,
		},
	}
	key := []byte(os.Getenv("SECRET_JWT_KEY"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(key)
}

func GetUserId(auth string) (string, error) {
	claims := &payload.ClaimsCustom{}
	splitToken := strings.Split(auth, "Bearer ")
	auth = splitToken[1]

	_, err := jwt.ParseWithClaims(auth, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_JWT_KEY")), nil
	})

	if err != nil {
		return "", errors.New("token has expired")
	}
	return claims.UserId, nil
}
