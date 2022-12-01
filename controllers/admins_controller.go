package controllers

import (
	"net/http"
	"vaksin-id-be/dto/payload"
	service_ad "vaksin-id-be/services/admins"

	"github.com/labstack/echo/v4"
)

type AdminsController struct {
	AdminServ service_ad.AdminService
}

func NewAdminsController(adminServ service_ad.AdminService) *AdminsController {
	return &AdminsController{
		AdminServ: adminServ,
	}
}

func (a *AdminsController) RegisterAdmin(ctx echo.Context) error {
	var payloads payload.AdminsPayload

	if err := ctx.Bind(&payloads); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	err := a.AdminServ.RegisterAdmin(payloads)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusCreated, map[string]interface{}{
		"error":    false,
		"messages": "success register admin",
	})
}

func (a *AdminsController) LoginAdmin(ctx echo.Context) error {
	var payloads payload.LoginAdmin

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

func (a *AdminsController) GetAllAdmin(ctx echo.Context) error {
	allData, err := a.AdminServ.GetAllAdmin()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"error":   false,
		"message": "success get all admins",
		"data":    allData,
	})
}

func (a *AdminsController) GetAdmin(ctx echo.Context) error {
	id := ctx.Param("id")

	admin, err := a.AdminServ.GetAdmin(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"error":   false,
		"message": "success get admin",
		"data":    admin,
	})
}

func (a *AdminsController) DeleteAdmin(ctx echo.Context) error {
	id := ctx.Param("id")

	if err := a.AdminServ.DeleteAdmin(id); err != nil {
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

func (a *AdminsController) UpdateAdmin(ctx echo.Context) error {
	var payloads payload.AdminsPayload

	if err := ctx.Bind(&payloads); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	id := ctx.Param("id")

	if err := a.AdminServ.UpdateAdmin(payloads, id); err != nil {
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
