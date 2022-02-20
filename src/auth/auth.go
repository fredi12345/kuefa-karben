package auth

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

type (
	Provider interface {
		Login(c echo.Context) error
		Callback(c echo.Context) error
		Logout(c echo.Context) error
	}

	Context struct {
		echo.Context
		User string
	}
)

func NewProvider() (Provider, error) {
	provider := viper.GetString("auth.provider")
	switch provider {
	case "local":
		return &localProvider{
			name:     viper.GetString("auth.session.name"),
			lifetime: viper.GetDuration("auth.session.lifetime"),
			secret:   []byte(viper.GetString("auth.session.secret")),
		}, nil
	}

	return nil, fmt.Errorf("provider with name '%s' not found", provider)
}
