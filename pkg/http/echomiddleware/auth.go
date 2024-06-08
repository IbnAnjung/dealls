package echomiddleware

import (
	"strings"

	coreerror "github.com/IbnAnjung/dealls/pkg/error"
	"github.com/IbnAnjung/dealls/pkg/http"
	"github.com/IbnAnjung/dealls/pkg/jwt"

	"github.com/labstack/echo/v4"
)

func AuthenticationMiddleware(jwtService jwt.JwtService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			bearer := c.Request().Header.Get("Authorization")
			tokens := strings.Split(bearer, " ")

			if bearer == "" || len(tokens) != 2 {
				err := coreerror.NewCoreError(coreerror.CoreErrorTypeAuthorization, "")
				return err
			}

			claim, err := jwtService.ValidateToken(tokens[1])
			if err != nil {
				err := coreerror.NewCoreError(coreerror.CoreErrorTypeAuthorization, err.Error())
				return err
			}

			c.Set(http.UserIdContextKey, claim.UserID)
			c.Set(http.UserIsPremiumContextKey, claim.IsPremium)

			return next(c)
		}
	}
}
