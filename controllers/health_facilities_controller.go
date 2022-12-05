package controllers

import (
	"net/http"
	"strings"
	"vaksin-id-be/dto/payload"
	service_a "vaksin-id-be/services/addresses"
	service_h "vaksin-id-be/services/health_facilities"
	"vaksin-id-be/util"

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

// @Summary 	Create HealthFacilities
// @Description Create data for Health Facilities
// @Tags 		HealthFacilities
// @Accept		json
// @Produce 	json
// @Param		healthfacilities body 	payload.HealthFacilities	true	"Input data Health Facilities"
// @Success 	201		{object} 	response.Response{data=payload.HealthFacilities}		"success create health facilities"
// @Router 		/api/v1/admin/healthfacilities [post]
// @failure		400		{object}	response.ResponseError	"StatusBadRequest"
// @failure		500		{object}	response.ResponseError	"StatusInternalServerError"
func (h *HealthFacilitiesController) CreateHealthFacilities(ctx echo.Context) error {
	var payloads payload.HealthFacilities

	if err := ctx.Bind(&payloads); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	if err := util.ValidateHealthFacilities(payloads); err != nil {
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

// @Summary		Get HealthFacilities by Name
// @Description This can only be done by the logged in admin.
// @Tags		HealthFacilities
// @Produce		json
// @Success		200	{object}	response.Response{data=response.HealthFacilitiesSwagger}	"Success get health facilities"
// @Router		/api/v1/healthfacilities/:name [get]
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

// @Summary 	Update HealthFacilities
// @Description This can only be done by the logged in admin.
// @Tags 		HealthFacilities
// @Accept		json
// @Produce 	json
// @Param		update body 	payload.UpdateHealthFacilities	true	"Input new data health facilities"
// @Success 	200		{object} 	response.Response{data=payload.HealthFacilities}		"success update health facilities"
// @Router 		/api/v1/admin/healthfacilities/:id [put]
// @failure		400		{object}		response.ResponseError	"StatusBadRequest"
func (h *HealthFacilitiesController) UpdateHealthFacilities(ctx echo.Context) error {
	var payloads payload.UpdateHealthFacilities

	id := ctx.Request().Header.Get("Authorization")

	if err := ctx.Bind(&payloads); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	if err := h.HealthService.UpdateHealthFacilities(payloads, id); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"error":   false,
		"message": "success update health facilities",
	})
}

// @Summary 	Delete HealthFacilities
// @Description delete data healthfacilities
// @Tags 		HealthFacilities
// @Produce 	json
// @Param       uuid   path      string  true  "Account ID"
// @Success 	200		{object} 	response.ResponseDelete	"success delete healthfacilities"
// @Router 		/api/v1/admin/healthfacilities/:id [delete]
// @failure		401		{object}		response.ResponseError	"StatusUnauthorized"
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
