package controllers

import (
	"net/http"
	"vaksin-id-be/dto/payload"
	_ "vaksin-id-be/dto/response"
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

// @Summary 	Register Users
// @Description Register User
// @Tags 		users
// @Accept		json
// @Produce 	json
// @Param		register body 	payload.RegisterUser	true	"Input Register User"
// @Success 	201		{object} 	response.Response{data=payload.RegisterUser}		"success register user"
// @Router 		/api/v1/signup [post]
// @failure		400		{string}	string	"error"
func (u *UserController) RegisterUser(ctx echo.Context) error {
	var payloads payload.RegisterUser

	if err := ctx.Bind(&payloads); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	if err := util.ValidateRegister(payloads); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	err := u.UserService.RegisterUser(payloads)

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

// @Summary 	Login Users
// @Description Login User
// @Tags 		users
// @Accept		json
// @Produce 	json
// @Param		login body 	payload.Login	true	"Input Login User"
// @Success 	200		{object} 	response.Response{data=response.Login}		"success login user"
// @Router 		/api/v1/login [post]
// @failure		400		{string}	string	"error"
// @failure		401		{string}	string	"error"
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
		"error":   false,
		"message": "succes login user",
		"data":    authUser,
	})
}

// @Summary		get user by nik
// @Description	This can only be done by the logged in user.
// @Tags		users
// @Produce		json
// @Success		200	{object}	response.Response{data=response.UserProfile}	"Success get user"
// @Router		/api/v1/profile [get]
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
		"error":   false,
		"message": "succes get user",
		"data":    data,
	})
}

// @Summary 	update users
// @Description This can only be done by the logged in user.
// @Tags 		users
// @Accept		json
// @Produce 	json
// @Param		update body 	payload.UpdateUser	true	"Input new data user"
// @Success 	200		{object} 	response.Response{data=payload.UpdateUser}		"success update user"
// @Router 		/api/v1/profile [put]
// @failure		400		{string}	string	"error"
// @failure		401		{string}	string	"error"
func (u *UserController) UpdateUser(ctx echo.Context) error {
	var payloads payload.UpdateUser

	if err := ctx.Bind(&payloads); err != nil {
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
		"error":   false,
		"message": "succes update user",
	})
}

// @Summary		get user address
// @Description	This can only be done by the logged in user.
// @Tags		users
// @Produce		json
// @Success		200	{object}	response.Response{data=response.UserAddresses}	"Success get user address"
// @Router		/api/v1/profile/address [get]
func (u *UserController) GetUserAddress(ctx echo.Context) error {
	nik := ctx.Request().Header.Get("Authorization")
	data, err := u.UserService.GetAddressUser(nik)

	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"error":   false,
		"message": "succes get user address",
		"data":    data,
	})
}

// @Summary 	delete user
// @Description delete data users
// @Tags 		users
// @Produce 	json
// @Success 	200		{string} 	string	"success delete user"
// @Router 		/api/v1/profile [delete]
// @failure		401		{string}	string	"error"
func (u *UserController) DeleteUser(ctx echo.Context) error {
	nik := ctx.Request().Header.Get("Authorization")

	if err := u.UserService.DeleteUserProfile(nik); err != nil {
		return ctx.JSON(http.StatusUnauthorized, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"error":   false,
		"message": "succes delete user",
	})
}

// @Summary 	update users addres
// @Description This can only be done by the logged in user.
// @Tags 		users
// @Accept		json
// @Produce 	json
// @Param		update body 	payload.UpdateAddress	true	"Input new data user address"
// @Success 	200		{object} 	response.Response{data=payload.UpdateUser}		"success update address user"
// @Router 		/api/v1/profile/address [put]
// @failure		400		{string}	string	"error"
// @failure		401		{string}	string	"error"
func (u *UserController) UpdateUserAddress(ctx echo.Context) error {
	var payloads payload.UpdateAddress

	if err := ctx.Bind(&payloads); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	nik := ctx.Request().Header.Get("Authorization")

	if err := u.UserService.UpdateUserAddress(payloads, nik); err != nil {
		return ctx.JSON(http.StatusUnauthorized, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"error":   false,
		"message": "succes update address user",
	})
}
