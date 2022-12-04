package controllers

import (
	"net/http"
	"vaksin-id-be/dto/payload"
	service_a "vaksin-id-be/services/admins"

	"github.com/labstack/echo/v4"
)

type AdminController struct {
	AdminServ service_a.AdminService
}

func NewAdminController(adminServ service_a.AdminService) *AdminController {
	return &AdminController{
		AdminServ: adminServ,
	}
}

func (a *AdminController) LoginAdmin(ctx echo.Context) error {
	var payloads payload.Login

	if err := ctx.Bind(&payloads); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	authAdmin, err := a.AdminServ.LoginAdmin(payloads)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"error":   false,
		"message": "success login admin",
		"data":    authAdmin,
	})
}

func (a *AdminController) GetAdmins(ctx echo.Context) error {
	id := ctx.Param("id")
	data, err := a.AdminServ.GetAdmins(id)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"error":   false,
		"message": "success get admin",
		"data":    data,
	})
}

func (a *AdminController) GetAllAdmins(ctx echo.Context) error {
	data, err := a.AdminServ.GetAllAdmins()

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"error":   false,
		"message": "success get all admin",
		"data":    data,
	})
}

func (a *AdminController) UpdateAdmins(ctx echo.Context) error {
	var payloads payload.AdminsPayload

	id := ctx.Param("id")

	if err := ctx.Bind(&payloads); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	if err := a.AdminServ.UpdateAdmins(payloads, id); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"error":   false,
		"message": "success update admin",
	})
}

func (a *AdminController) DeleteAdmins(ctx echo.Context) error {
	id := ctx.Param("id")

	if err := a.AdminServ.DeleteAdmins(id); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"error":   false,
		"message": "success delete admin",
	})
}
