package middleware

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type ClaimsCustom struct {
	NikUser string
	Email   string
	jwt.StandardClaims
}
type ClaimsCustomAdmin struct {
	Id                 string
	IdHealthFacilities string
	Email              string
	jwt.StandardClaims
}

func CreateToken(nikUser string, email string) (string, error) {
	exp := time.Now().Add(time.Hour * 72).Unix()
	claims := &ClaimsCustom{
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

func CreateTokenAdmin(id, idhealthfacil, email string) (string, error) {
	exp := time.Now().Add(time.Hour * 72).Unix()
	claims := &ClaimsCustomAdmin{
		Id:                 id,
		IdHealthFacilities: idhealthfacil,
		Email:              email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp,
		},
	}
	key := []byte(os.Getenv("SECRET_JWT_KEY_ADMIN"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(key)
}
