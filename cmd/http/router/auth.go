package router

import (
	"github.com/IbnAnjung/dealls/cmd/http/handler"
	"github.com/IbnAnjung/dealls/entity/enauth"

	"github.com/labstack/echo/v4"
)

func LoadAuthRouter(e *echo.Echo, authUC enauth.AuthUsecase) {
	h := handler.NewAuthHandler(authUC)

	ur := e.Group("/auth")
	ur.POST("/register", h.Register)
	ur.POST("/login", h.Login)
}
