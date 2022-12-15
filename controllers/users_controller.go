package controllers

import (
	"net/http"
	"vaksin-id-be/dto/payload"
	service_a "vaksin-id-be/services/addresses"
	service_u "vaksin-id-be/services/users"
	"vaksin-id-be/util"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	UserService    service_u.UserService
	AddressService service_a.AddressService
}

func NewUserController(userServ service_u.UserService, addressServ service_a.AddressService) *UserController {
	return &UserController{
		UserService:    userServ,
		AddressService: addressServ,
	}
}

// @Summary 	Register Users
// @Description Register User
// @Tags 		Authentication
// @Accept		json
// @Produce 	json
// @Param		register body 	payload.RegisterUser	true	"Input Register User"
// @Success 	201		{object} 	response.Response{data=payload.RegisterUser}		"success register user"
// @Router 		/api/v1/signup [post]
// @failure		400		{object}	response.ResponseError	"StatusBadRequest"
// @failure		500		{object}	response.ResponseError	"StatusInternalServerError"
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
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
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
// @Tags 		Authentication
// @Accept		json
// @Produce 	json
// @Param		login body 	payload.Login	true	"Input Login User"
// @Success 	200		{object} 	response.Response{data=response.Login}		"success login user"
// @Router 		/api/v1/login [post]
// @failure		400		{object}		response.ResponseError	"StatusBadRequest"
// @failure		401		{object}	response.ResponseError	"StatusUnauthorized"
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
		"message": "success login user",
		"data":    authUser,
	})
}

// @Summary		Get User by NIK
// @Description	This can only be done by the logged in user.
// @Tags		Users
// @Produce		json
// @Success		200	{object}	response.Response{data=response.UserProfile}	"Success get user"
// @Router		/api/v1/profile [get]
func (u *UserController) GetUserDataByNik(ctx echo.Context) error {
	nik := ctx.Request().Header.Get("Authorization")
	data, err := u.UserService.GetUserDataByNik(nik)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"error":   false,
		"message": "success get user",
		"data":    data,
	})
}

func (u *UserController) GetUserDataByNikCheck(ctx echo.Context) error {
	nik := ctx.Param("nik")
	data, err := u.UserService.GetUserDataByNikNoAddress(nik)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"error":   false,
		"message": "success check user nik",
		"data":    data,
	})
}

// @Summary 	Update User
// @Description This can only be done by the logged in user.
// @Tags 		Users
// @Accept		json
// @Produce 	json
// @Param		update body 	payload.UpdateUser	true	"Input new data user"
// @Success 	200		{object} 	response.Response{data=payload.UpdateUser}		"success update user"
// @Router 		/api/v1/profile [put]
// @failure		400		{object}		response.ResponseError	"StatusBadRequest"
func (u *UserController) UpdateUser(ctx echo.Context) error {
	var payloads payload.UpdateUser

	if err := ctx.Bind(&payloads); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	nik := ctx.Request().Header.Get("Authorization")

	data, err := u.UserService.UpdateUserProfile(payloads, nik)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"error":   false,
		"message": "success update user",
		"data":    data,
	})
}

// @Summary		Get User Address
// @Description	This can only be done by the logged in user.
// @Tags		Users
// @Produce		json
// @Success		200	{object}	response.Response{data=response.UserAddresses}	"Success get user address"
// @Router		/api/v1/profile/address [get]
func (u *UserController) GetUserAddress(ctx echo.Context) error {
	nik := ctx.Request().Header.Get("Authorization")
	data, err := u.AddressService.GetAddressUser(nik)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"error":   false,
		"message": "success get user address",
		"data":    data,
	})
}

// @Summary 	Delete User
// @Description This can only be done by the logged in user.
// @Tags 		Users
// @Produce 	json
// @Success 	200		{object} 	response.ResponseDelete	"success delete user"
// @Router 		/api/v1/profile [delete]
// @failure		401		{object}		response.ResponseError	"StatusUnauthorized"
func (u *UserController) DeleteUser(ctx echo.Context) error {
	nik := ctx.Request().Header.Get("Authorization")

	if err := u.UserService.DeleteUserProfile(nik); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"error":   false,
		"message": "success delete user",
	})
}

// @Summary 	Update User Address
// @Description This can only be done by the logged in user.
// @Tags 		Users
// @Accept		json
// @Produce 	json
// @Param		update body 	payload.UpdateAddress	true	"Input new data user address"
// @Success 	200		{object} 	response.Response{data=payload.UpdateUser}		"success update address user"
// @Router 		/api/v1/profile/address [put]
// @failure		400		{object}		response.ResponseError	"StatusBadRequest"
// @failure		401		{object}	response.ResponseError	"StatusUnauthorized"
func (u *UserController) UpdateUserAddress(ctx echo.Context) error {
	var payloads payload.UpdateAddress

	if err := ctx.Bind(&payloads); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	nik := ctx.Request().Header.Get("Authorization")

	data, err := u.AddressService.UpdateUserAddress(payloads, nik)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"error":   false,
		"message": "success update address user",
		"data":    data,
	})
}

func (u *UserController) UserNearbyHealth(ctx echo.Context) error {
	nik := ctx.Request().Header.Get("Authorization")
	var payloads payload.NearbyHealth

	if err := ctx.Bind(&payloads); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	data, err := u.UserService.NearbyHealthFacilities(payloads, nik)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"error":   false,
		"message": "success get nearby health facilities",
		"data":    data,
	})
}
