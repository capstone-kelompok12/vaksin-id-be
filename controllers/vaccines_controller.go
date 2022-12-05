package controllers

import (
	"net/http"
	"vaksin-id-be/dto/payload"
	service_v "vaksin-id-be/services/vaccines"
	"vaksin-id-be/util"

	"github.com/labstack/echo/v4"
)

type VaccinesController struct {
	VaccineService service_v.VaccinesService
}

func NewVaccinesController(vaccineServ service_v.VaccinesService) *VaccinesController {
	return &VaccinesController{
		VaccineService: vaccineServ,
	}
}

func (v *VaccinesController) CreateVaccine(ctx echo.Context) error {
	var payloads payload.VaccinesPayload

	if err := ctx.Bind(&payloads); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	if err := util.ValidateVaccine(payloads); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	authAdmin := ctx.Request().Header.Get("Authorization")

	if err := v.VaccineService.CreateVaccine(authAdmin, payloads); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusCreated, map[string]interface{}{
		"error":    false,
		"messages": "success create vaccine by admin",
	})
}

func (v *VaccinesController) GetAllVaccines(ctx echo.Context) error {
	dataAllVaccine, err := v.VaccineService.GetAllVaccines()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"error":    false,
		"messages": "success get all vaccines",
		"data":     dataAllVaccine,
	})
}

func (v *VaccinesController) GetVaccineByAdmin(ctx echo.Context) error {
	authAdmin := ctx.Request().Header.Get("Authorization")

	dataVaccines, err := v.VaccineService.GetVaccineByAdmin(authAdmin)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"error":    false,
		"messages": "success get all vaccines by admin",
		"data":     dataVaccines,
	})
}

func (v *VaccinesController) UpdateVaccine(ctx echo.Context) error {
	var payloads payload.VaccinesUpdatePayload

	if err := ctx.Bind(&payloads); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	id := ctx.Param("id")

	if err := v.VaccineService.UpdateVaccine(id, payloads); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"error":    false,
		"messages": "success update vaccine by admin",
	})
}

func (v *VaccinesController) DeleteVacccine(ctx echo.Context) error {
	id := ctx.Param("id")

	if err := v.VaccineService.DeleteVacccine(id); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"error":   false,
		"message": "success delete vaccine by admin",
	})
}
