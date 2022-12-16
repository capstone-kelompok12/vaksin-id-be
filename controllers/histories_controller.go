package controllers

import (
	"net/http"
	"vaksin-id-be/dto/payload"
	service_h "vaksin-id-be/services/histories"

	"github.com/labstack/echo/v4"
)

type HistoriesController struct {
	HistoriesService service_h.HistoriesRepository
}

func NewHistoriesController(historyServ service_h.HistoriesRepository) *HistoriesController {
	return &HistoriesController{
		HistoriesService: historyServ,
	}
}

func (h *HistoriesController) CreateHistory(ctx echo.Context) error {
	var payloads payload.HistoriesPayload

	if err := ctx.Bind(&payloads); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	err := h.HistoriesService.CreateHistory(payloads)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusCreated, map[string]interface{}{
		"error":    false,
		"messages": "success create health history",
	})
}

func (h *HistoriesController) GetAllHistory(ctx echo.Context) error {
	allData, err := h.HistoriesService.GetAllHistory()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"error":    false,
		"messages": "success get all data history",
		"data":     allData,
	})
}

func (h *HistoriesController) GetHistoryById(ctx echo.Context) error {
	id := ctx.Param("id")

	data, err := h.HistoriesService.GetHistoryById(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"error":    false,
		"messages": "success get data history",
		"data":     data,
	})
}

func (h *HistoriesController) UpdateHistory(ctx echo.Context) error {
	var payloads payload.UpdateAccHistory

	id := ctx.Request().Header.Get("Authorization")

	if err := ctx.Bind(&payloads); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	data, err := h.HistoriesService.UpdateHistory(payloads, id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"error":   false,
		"message": "success update history",
		"data":    data,
	})
}
