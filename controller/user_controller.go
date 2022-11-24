package controller

import (
	"net/http"
	"vaksin-id-be/dto/payload"
	service "vaksin-id-be/service/user"
	"vaksin-id-be/util"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(userServ service.UserService) *UserController {
	return &UserController{
		UserService: userServ,
	}
}

func (u *UserController) RegisterUser(ctx echo.Context) error {
	var payload payload.Register

	if err := ctx.Bind(&payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	if err := util.ValidateRegister(payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	err := u.UserService.RegisterUser(payload)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusCreated, map[string]interface{}{
		"messages": "success create user",
	})
}

func (u *UserController) LoginUser(ctx echo.Context) error {
	var payload payload.Login

	if err := ctx.Bind(&payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": err.Error(),
		})
	}

	authUser, err := u.UserService.LoginUser(payload)

	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"user token": authUser,
	})
}
