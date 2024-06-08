package handler

import (
	"net/http"

	"github.com/IbnAnjung/dealls/entity/enauth"
	coreerror "github.com/IbnAnjung/dealls/pkg/error"
	pkgHttp "github.com/IbnAnjung/dealls/pkg/http"

	"github.com/IbnAnjung/dealls/cmd/http/handler/presenter"

	"github.com/labstack/echo/v4"
)

type authHandler struct {
	uc enauth.AuthUsecase
}

func NewAuthHandler(
	uc enauth.AuthUsecase,
) *authHandler {
	return &authHandler{
		uc,
	}
}

func (h authHandler) Register(c echo.Context) error {
	req := presenter.RegisterRequest{}
	requestdId := c.Get(pkgHttp.RequestIdContextKey).(string)

	if err := c.Bind(&req); err != nil {
		c.Logger().Errorf("fail binding request request_id:%s, %v", requestdId, err)
		e := coreerror.NewCoreError(coreerror.CoreErrorTypeInternalServerError, "")
		err = e
		return err
	}

	output, err := h.uc.RegisterUser(c.Request().Context(), enauth.RegisterInput{
		Username:        req.Username,
		Password:        req.Password,
		Fullname:        req.Fullname,
		Gender:          req.Gender,
		Age:             req.Age,
		ConfirmPassword: req.ConfirmPassword,
	})
	if err != nil {
		c.Logger().Errorf("fail uc - request request_id:%s, %v", requestdId, err)
		return err
	}

	res := presenter.RegisterResponse{
		ID:           output.ID,
		Fullname:     output.Fullname,
		AccessToken:  output.AccessToken,
		RefreshToken: output.RefreshToken,
	}

	return c.JSON(http.StatusCreated, pkgHttp.GetStandartSuccessResponse("success", res))
}

func (h authHandler) Login(c echo.Context) error {
	req := presenter.LoginRequest{}
	requestdId := c.Get(pkgHttp.RequestIdContextKey).(string)

	if err := c.Bind(&req); err != nil {
		e := coreerror.NewCoreError(coreerror.CoreErrorTypeInternalServerError, "")
		err = e
		c.Logger().Errorf("fail binding request request_id:%s, %v", requestdId, err)
		return err
	}

	output, err := h.uc.Login(c.Request().Context(), enauth.LoginInput{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		c.Logger().Errorf("fail uc - request request_id:%s, %v", requestdId, err)
		return err
	}

	res := presenter.LoginResponse{
		ID:           output.ID,
		Fullname:     output.Fullname,
		IsPremium:    output.IsPremium,
		AccessToken:  output.AccessToken,
		RefreshToken: output.RefreshToken,
	}

	return c.JSON(http.StatusOK, pkgHttp.GetStandartSuccessResponse("success", res))
}
