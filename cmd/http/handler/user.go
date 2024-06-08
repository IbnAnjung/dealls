package handler

import (
	"net/http"

	"github.com/IbnAnjung/dealls/entity/enuser"
	pkgHttp "github.com/IbnAnjung/dealls/pkg/http"

	"github.com/IbnAnjung/dealls/cmd/http/handler/presenter"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	uc enuser.UserUsecase
}

func NewUserHandler(
	uc enuser.UserUsecase,
) *userHandler {
	return &userHandler{
		uc,
	}
}

func (h userHandler) GetDatingProfile(c echo.Context) error {
	requestdId := c.Get(pkgHttp.RequestIdContextKey).(string)

	userID := c.Get(pkgHttp.UserIdContextKey).(string)

	output, err := h.uc.GetDatingUserProfile(c.Request().Context(), userID)
	if err != nil {
		c.Logger().Errorf("fail uc - request request_id:%s, %v", requestdId, err)
		return err
	}

	res := presenter.GetDatingUserProfileResponse{
		ID:       output.ID,
		Fullname: output.Fullname,
		Gender:   uint8(output.Gender),
		Age:      output.Age,
	}

	return c.JSON(http.StatusOK, pkgHttp.GetStandartSuccessResponse("success", res))
}

func (h userHandler) Upgrade(c echo.Context) error {
	requestdId := c.Get(pkgHttp.RequestIdContextKey).(string)

	userID := c.Get(pkgHttp.UserIdContextKey).(string)

	err := h.uc.UpgradeToPremium(c.Request().Context(), userID)
	if err != nil {
		c.Logger().Errorf("fail uc - request request_id:%s, %v", requestdId, err)
		return err
	}

	return c.JSON(http.StatusOK, pkgHttp.GetStandartSuccessResponse("success", nil))
}

func (h userHandler) Swipe(c echo.Context) error {
	requestdId := c.Get(pkgHttp.RequestIdContextKey).(string)
	req := presenter.SwipeUserRequest{}
	if err := c.Bind(&req); err != nil {
		c.Logger().Errorf("fail to bind request, request_id: %s, %v", requestdId, err)
		return err
	}

	userID := c.Get(pkgHttp.UserIdContextKey).(string)

	err := h.uc.Swipe(c.Request().Context(), enuser.SwipeUserInput{
		UserID:        userID,
		SwippedUserID: req.SwipeUserID,
		Type:          (*enuser.SwipeType)(req.Type),
	})
	if err != nil {
		c.Logger().Errorf("fail uc - request request_id:%s, %v", requestdId, err)
		return err
	}

	return c.JSON(http.StatusOK, pkgHttp.GetStandartSuccessResponse("success", nil))
}
