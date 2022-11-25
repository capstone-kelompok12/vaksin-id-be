package controllers

import (
	"net/http"
	"vaksin-id-be/dto/payload"
	service "vaksin-id-be/services/users"
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
	var payload payload.RegisterUser

	if err := ctx.Bind(&payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	if err := util.Validate(payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	err := u.UserService.RegisterUser(payload)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusCreated, map[string]interface{}{
		"error":    false,
		"messages": "success register user",
	})
}

func (u *UserController) LoginUser(ctx echo.Context) error {
	var payload payload.Login

	if err := ctx.Bind(&payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	authUser, err := u.UserService.LoginUser(payload)

	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"error":   true,
		"message": "succes login user",
		"data": map[string]interface{}{
			"usertoken": authUser,
		},
	})
}

func (u *UserController) GetUserDataByNik(ctx echo.Context) error {
	nik := ctx.Request().Header.Get("Authorization")
	data, err := u.UserService.GetUserDataByNik(nik)

	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"error":   true,
		"message": "succes get user",
		"data":    data,
	})
}

func (u *UserController) UpdateUser(ctx echo.Context) error {
	var payloads payload.UpdateUser

	if err := ctx.Bind(&payloads); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	if err := util.Validate(payloads); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	nik := ctx.Request().Header.Get("Authorization")

	if err := u.UserService.UpdateUserProfile(payloads, nik); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"error":   true,
		"message": "succes update user",
	})
}

func (u *UserController) DeleteUser(ctx echo.Context) error {
	nik := ctx.Request().Header.Get("Authorization")

	if err := u.UserService.DeleteUserProfile(nik); err != nil {
		return ctx.JSON(http.StatusUnauthorized, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"error":   true,
		"message": "succes delete user",
	})
}
