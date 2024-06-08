package router

import (
	"github.com/IbnAnjung/dealls/cmd/http/handler"
	"github.com/IbnAnjung/dealls/entity/enuser"
	"github.com/IbnAnjung/dealls/pkg/http/echomiddleware"
	"github.com/IbnAnjung/dealls/pkg/jwt"

	"github.com/labstack/echo/v4"
)

func LoadUserRouter(e *echo.Echo, userUC enuser.UserUsecase, jwtService jwt.JwtService) {
	h := handler.NewUserHandler(userUC)

	authMiddleware := echomiddleware.AuthenticationMiddleware(jwtService)

	ur := e.Group("/user", authMiddleware)
	ur.GET("/datting-profile", h.GetDatingProfile)
	ur.PATCH("/upgrade", h.Upgrade)
	ur.POST("/swipe", h.Swipe)
}
