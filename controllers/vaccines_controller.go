package controllers

import (
	"net/http"
	"vaksin-id-be/dto/payload"
	service_v "vaksin-id-be/services/vaccines"
	"vaksin-id-be/util"

	"github.com/golang-jwt/jwt"
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

	admin := ctx.Get("user").(*jwt.Token)
	claimId := admin.Claims.(jwt.MapClaims)
	id := claimId["IdHealthFacilities"].(string)

	data, err := v.VaccineService.CreateVaccine(id, payloads)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusCreated, map[string]interface{}{
		"error":    false,
		"messages": "success create vaccine by admin",
		"data":     data,
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

func (v *VaccinesController) GetAllVaccinesCount(ctx echo.Context) error {
	data, err := v.VaccineService.GetVaccinesCount()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"error":    false,
		"messages": "success get all vaccines dose",
		"data":     data,
	})
}

func (v *VaccinesController) GetVaccineByAdmin(ctx echo.Context) error {
	admin := ctx.Get("user").(*jwt.Token)
	claimId := admin.Claims.(jwt.MapClaims)
	id := claimId["IdHealthFacilities"].(string)

	dataVaccines, err := v.VaccineService.GetVaccineByAdmin(id)
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

	data, err := v.VaccineService.UpdateVaccine(id, payloads)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"error":    false,
		"messages": "success update vaccine by admin",
		"data":     data,
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

func (v *VaccinesController) GetVaccineDashboard(ctx echo.Context) error {
	dataAllVaccine, err := v.VaccineService.GetVaccineDashboard()
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
