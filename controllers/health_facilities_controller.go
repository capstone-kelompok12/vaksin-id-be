package controllers

import (
	"net/http"
	"strings"
	"vaksin-id-be/dto/payload"
	service_a "vaksin-id-be/services/addresses"
	service_h "vaksin-id-be/services/health_facilities"
	"vaksin-id-be/util"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type HealthFacilitiesController struct {
	HealthService  service_h.HealthFacilitiesService
	AddressService service_a.AddressService
}

func NewHealthFacilitiesController(healthServ service_h.HealthFacilitiesService, addressServ service_a.AddressService) *HealthFacilitiesController {
	return &HealthFacilitiesController{
		HealthService:  healthServ,
		AddressService: addressServ,
	}
}

func (h *HealthFacilitiesController) CreateHealthFacilities(ctx echo.Context) error {
	var payloads payload.HealthFacilities

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

	err := h.HealthService.CreateHealthFacilities(payloads)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusCreated, map[string]interface{}{
		"error":    false,
		"messages": "success create health facilities",
	})
}

func (h *HealthFacilitiesController) GetHealthFacilities(ctx echo.Context) error {
	name := ctx.Param("name")
	nameLower := strings.ToLower(name)
	data, err := h.HealthService.GetHealthFacilities(nameLower)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"error":   false,
		"message": "success get health facilities",
		"data":    data,
	})
}

func (h *HealthFacilitiesController) GetAllHealthFacilities(ctx echo.Context) error {
	data, err := h.HealthService.GetAllHealthFacilities()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"error":   false,
		"message": "success get all health facilites",
		"data":    data,
	})
}

func (h *HealthFacilitiesController) UpdateHealthFacilities(ctx echo.Context) error {
	var payloads payload.UpdateHealthFacilities

	admin := ctx.Get("user").(*jwt.Token)
	claimId := admin.Claims.(jwt.MapClaims)
	id := claimId["IdHealthFacilities"].(string)

	if err := ctx.Bind(&payloads); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	data, err := h.HealthService.UpdateHealthFacilities(payloads, id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"error":   false,
		"message": "success update health facilities",
		"data":    data,
	})
}

func (h *HealthFacilitiesController) DeleteHealthFacilities(ctx echo.Context) error {
	id := ctx.Param("id")

	if err := h.HealthService.DeleteHealthFacilities(id); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"error":   false,
		"message": "success delete health facilities",
	})
}
