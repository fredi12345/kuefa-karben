package auth

import (
	"fmt"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	name := viper.GetString("auth.session.secret")
	secret := []byte(viper.GetString("auth.session.secret"))

	return func(c echo.Context) error {
		cookie, err := c.Cookie(name)
		if err != nil {
			c.Logger().Infof("request without authentication cookie: %v", err)
			return echo.ErrUnauthorized
		}

		token, err := jwt.ParseWithClaims(cookie.Value, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method %v", token.Header["alg"])
			}

			return secret, nil
		})
		if err != nil {
			c.Logger().Infof("failed to parse token in authentication cookie: %v", err)
			return echo.ErrForbidden
		}

		claims, ok := token.Claims.(*jwt.StandardClaims)
		if !ok || !token.Valid {
			c.Logger().Warnf("invalid token claims for token %s", token.Raw)
			return echo.ErrInternalServerError
		}

		return next(&Context{
			Context: c,
			User:    claims.Subject,
		})
	}
}

func Authenticated(fn func(*Context) error) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := c.(*Context)
		return fn(cc)
	}
}
