package auth

import (
	_ "embed"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type (
	localProvider struct {
		name     string
		lifetime time.Duration
		secret   []byte
	}
)

//go:embed embed/login.html
var localLoginPage []byte

func (p *localProvider) Login(c echo.Context) error {
	return c.HTMLBlob(http.StatusOK, localLoginPage)
}

func (p *localProvider) Callback(c echo.Context) error {
	now := time.Now()
	user := c.QueryParam("user")

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.StandardClaims{
		ExpiresAt: now.Add(p.lifetime).Unix(),
		IssuedAt:  now.Unix(),
		Issuer:    p.name,
		Subject:   user,
	})

	signedToken, err := rawToken.SignedString(p.secret)
	if err != nil {
		c.Logger().Infof("failed to sign jwt: %v", err)
		return c.NoContent(http.StatusInternalServerError)
	}

	c.SetCookie(&http.Cookie{
		Name:     p.name,
		Path:     "/",
		Value:    signedToken,
		Expires:  now.Add(p.lifetime),
		Secure:   true,
		HttpOnly: false,
	})
	return c.Redirect(http.StatusFound, "/")
}

func (p *localProvider) Logout(c echo.Context) error {
	c.SetCookie(&http.Cookie{
		Name:    p.name,
		Path:    "/",
		Expires: time.Unix(0, 0),
	})
	return c.Redirect(http.StatusFound, "/")
}

var _ Provider = (*localProvider)(nil)
