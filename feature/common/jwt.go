package common

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		SigningKey:    []byte(os.Getenv("SECRET")),
	})
}

func GenerateToken(ID int, Role string) string {
	info := jwt.MapClaims{}
	info["ID"] = ID
	info["Role"] = Role
	info["exp"] = time.Now().Add(time.Minute * 180).Unix()
	auth := jwt.NewWithClaims(jwt.SigningMethodHS256, info)
	token, err := auth.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		log.Fatal("cannot generate key")
		return ""
	}

	return token
}

func ExtractData(c echo.Context) (int, string) {
	head := c.Request().Header
	token := strings.Split(head.Get("Authorization"), " ")

	res, _ := jwt.Parse(token[len(token)-1], func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET")), nil
	})

	if res.Valid {
		resClaim := res.Claims.(jwt.MapClaims)
		parseID := resClaim["ID"].(float64)
		parseRole := resClaim["Role"].(string)
		return int(parseID), parseRole
	}

	return -1, ""
}
