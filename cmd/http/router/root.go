package router

import (
	"github.com/IbnAnjung/dealls/entity/enauth"
	"github.com/IbnAnjung/dealls/entity/enuser"
	"github.com/IbnAnjung/dealls/pkg/jwt"
	"github.com/labstack/echo/v4"
)

func SetupRouter(
	e *echo.Echo,
	authUc enauth.AuthUsecase,
	userUc enuser.UserUsecase,
	jwtService jwt.JwtService,
) {

	LoadHealtRouter(e)
	LoadAuthRouter(e, authUc)
	LoadUserRouter(e, userUc, jwtService)
}
