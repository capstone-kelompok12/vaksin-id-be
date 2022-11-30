package middleware

import (
	"errors"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

type ClaimsCustom struct {
	NikUser string
	Email   string
	jwt.StandardClaims
}
type ClaimsCustomAdmin struct {
	Id    string
	Email string
	jwt.StandardClaims
}

func CreateToken(nikUser string, email string) (string, error) {
	exp := time.Now().Add(time.Hour * 1).Unix()
	claims := ClaimsCustom{
		NikUser: nikUser,
		Email:   email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp,
		},
	}
	key := []byte(os.Getenv("SECRET_JWT_KEY"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(key)
}

func GetUserNik(auth string) (string, error) {
	claims := &ClaimsCustom{}
	splitToken := strings.Split(auth, "Bearer ")
	auth = splitToken[1]

	_, err := jwt.ParseWithClaims(auth, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_JWT_KEY")), nil
	})

	if err != nil {
		return "", errors.New("token has expired")
	}
	return claims.NikUser, nil
}

func CreateTokenAdmin(id string, email string) (string, error) {
	exp := time.Now().Add(time.Hour * 1).Unix()
	claims := ClaimsCustomAdmin{
		Id:    id,
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp,
		},
	}
	key := []byte(os.Getenv("SECRET_JWT_KEY_ADMIN"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(key)
}

func GetIdAdmin(auth string) (string, error) {
	claims := &ClaimsCustomAdmin{}
	splitToken := strings.Split(auth, "Bearer ")
	auth = splitToken[1]

	_, err := jwt.ParseWithClaims(auth, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_JWT_KEY_ADMIN")), nil
	})

	if err != nil {
		return "", errors.New("token has expired")
	}
	return claims.Id, nil
}
