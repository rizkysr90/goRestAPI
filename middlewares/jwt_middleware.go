package middlewares

import (
	constants "project/constant"
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtClaims struct {
	userId int
	jwt.StandardClaims
}

func GenerateTokenJWT(id int) (string, error) {
	claims := JwtClaims{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * 1).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtToken, err := token.SignedString([]byte(constants.JWT_SECRET))

	if err != nil {
		return "", err
	}

	return jwtToken, nil
}
