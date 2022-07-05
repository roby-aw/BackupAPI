package middleware

import (
	"errors"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func CustomerSetupAuthenticationJWT() echo.MiddlewareFunc {
	config := ConfigMiddleware("Customer")
	return middleware.JWTWithConfig(config)
}

func AdminSetupAuthenticationJWT() echo.MiddlewareFunc {
	config := ConfigMiddleware("Admin")
	return middleware.JWTWithConfig(config)
}

func StoreSetupAuthenticationJWT() echo.MiddlewareFunc {
	config := ConfigMiddleware("Store")
	return middleware.JWTWithConfig(config)
}

func ConfigMiddleware(fitur string) middleware.JWTConfig {
	SECRET_KEY := os.Getenv("SECRET_JWT")
	config := middleware.JWTConfig{
		ParseTokenFunc: func(auth string, c echo.Context) (interface{}, error) {
			keyFunc := func(t *jwt.Token) (interface{}, error) {
				if t.Method.Alg() != "HS256" {
					return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
				}
				return SECRET_KEY, nil
			}
			token, _ := jwt.Parse(auth, keyFunc)
			claims, _ := token.Claims.(jwt.MapClaims)
			if claims[fitur] == nil {
				return nil, errors.New("Role not " + fitur)
			}
			return token, nil
		},
	}
	return config
}
