package controllers

import (
	"net/http"
	"vaksin-id-be/dto/payload"
	service_a "vaksin-id-be/services/addresses"
	service_u "vaksin-id-be/services/users"
	"vaksin-id-be/util"

	"github.com/golang-jwt/jwt"
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

func (u *UserController) RegisterUser(ctx echo.Context) error {
	var payloads payload.RegisterUser

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

func (u *UserController) GetUserDataByNik(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.Token)
	claimId := user.Claims.(jwt.MapClaims)
	nik := claimId["NikUser"].(string)
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

func (u *UserController) GetUserRegisteredDashboard(ctx echo.Context) error {
	data, err := u.UserService.GetUserRegisteredDashboard()

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"error":   false,
		"message": "success get user registered statistics",
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

func (u *UserController) GetUserHistoryByNikCheck(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.Token)
	claimId := user.Claims.(jwt.MapClaims)
	nik := claimId["NikUser"].(string)

	data, err := u.UserService.GetUserHistory(nik)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"error":   false,
		"message": "success check user history",
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

	user := ctx.Get("user").(*jwt.Token)
	claimId := user.Claims.(jwt.MapClaims)
	nik := claimId["NikUser"].(string)

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

func (u *UserController) GetUserAddress(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.Token)
	claimId := user.Claims.(jwt.MapClaims)
	nik := claimId["NikUser"].(string)
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

func (u *UserController) DeleteUser(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.Token)
	claimId := user.Claims.(jwt.MapClaims)
	nik := claimId["NikUser"].(string)

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

func (u *UserController) UpdateUserAddress(ctx echo.Context) error {
	var payloads payload.UpdateAddress

	if err := ctx.Bind(&payloads); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	user := ctx.Get("user").(*jwt.Token)
	claimId := user.Claims.(jwt.MapClaims)
	nik := claimId["NikUser"].(string)

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
	user := ctx.Get("user").(*jwt.Token)
	claimId := user.Claims.(jwt.MapClaims)
	nik := claimId["NikUser"].(string)
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
