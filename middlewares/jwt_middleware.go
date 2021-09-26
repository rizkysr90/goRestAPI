package middlewares

import (
	"errors"
	constants "project/constant"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type JwtClaims struct {
	userId int
	jwt.StandardClaims
}

func GenerateTokenJWTUser(id int) (string, error) {
	claims := JwtClaims{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * 1).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtToken, err := token.SignedString([]byte(constants.JWT_SECRET_USER))

	if err != nil {
		return "", err
	}

	return jwtToken, nil
}
func GenerateTokenJWTAdmin(id int) (string, error) {
	claims := JwtClaims{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * 1).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtToken, err := token.SignedString([]byte(constants.JWT_SECRET_ADMIN))

	if err != nil {
		return "", err
	}

	return jwtToken, nil
}
func GetClaimsUser(c echo.Context) (int, error) {
	user := c.Get("user")
	if user != nil {
		userJwt := user.(*jwt.Token)
		if userJwt.Valid {
			claims := userJwt.Claims.(jwt.MapClaims)
			userId := claims["userId"].(float64)
			return int(userId), nil
		}
	}
	return 0, errors.New("failed get claims data")
}
func GetClaimsAdmin(c echo.Context) (int, error) {
	admin := c.Get("admin")
	if admin != nil {
		adminJwt := admin.(*jwt.Token)
		if adminJwt.Valid {
			claims := adminJwt.Claims.(jwt.MapClaims)
			adminId := claims["adminId"].(float64)
			return int(adminId), nil
		}
	}
	return 0, errors.New("failed get claims data")
}
