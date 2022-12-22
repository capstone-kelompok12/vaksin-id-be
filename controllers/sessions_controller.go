package controllers

import (
	"net/http"
	"vaksin-id-be/dto/payload"
	service_s "vaksin-id-be/services/sessions"
	"vaksin-id-be/util"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type SessionsController struct {
	SessionService service_s.SessionsService
}

func NewSessionsController(sessionServ service_s.SessionsService) *SessionsController {
	return &SessionsController{
		SessionService: sessionServ,
	}
}

func (s *SessionsController) CreateSession(ctx echo.Context) error {
	var payloads payload.SessionsPayload

	if err := ctx.Bind(&payloads); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	if err := util.ValidateSession(payloads); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	data, err := s.SessionService.CreateSessions(payloads)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusCreated, map[string]interface{}{
		"error":    false,
		"messages": "success create sessions by admin",
		"data":     data,
	})
}

func (s *SessionsController) GetAllSessions(ctx echo.Context) error {
	data, err := s.SessionService.GetAllSessions()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"error":   false,
		"message": "success get all sessions",
		"data":    data,
	})
}

func (s *SessionsController) GetAllSessionsByAdmin(ctx echo.Context) error {
	admin := ctx.Get("user").(*jwt.Token)
	claimId := admin.Claims.(jwt.MapClaims)
	id := claimId["IdHealthFacilities"].(string)

	data, err := s.SessionService.GetAllSessionsByAdmin(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"error":   false,
		"message": "success get all sessions",
		"data":    data,
	})
}

func (s *SessionsController) GetAllFinishedSessionCount(ctx echo.Context) error {

	data, err := s.SessionService.GetAllFinishedSessionCount()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"error":   false,
		"message": "success get all amount finished sessions",
		"data":    data,
	})
}

func (s *SessionsController) GetSessionsById(ctx echo.Context) error {
	id := ctx.Param("id")

	data, err := s.SessionService.GetSessionsById(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"error":   false,
		"message": "success get session by id",
		"data":    data,
	})
}

func (s *SessionsController) UpdateSession(ctx echo.Context) error {
	var payloads payload.SessionsUpdate

	id := ctx.Param("id")

	if err := ctx.Bind(&payloads); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	data, err := s.SessionService.UpdateSession(payloads, id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"error":   false,
		"message": "success get session admin by id",
		"data":    data,
	})
}

func (s *SessionsController) IsCloseSession(ctx echo.Context) error {
	var payloads payload.SessionsIsClose
	id := ctx.Param("id")

	if err := ctx.Bind(&payloads); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	data, err := s.SessionService.IsCloseSession(payloads, id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"error":   false,
		"message": "success update session status",
		"data":    data,
	})
}

func (s *SessionsController) DeleteSession(ctx echo.Context) error {
	id := ctx.Param("id")

	if err := s.SessionService.DeleteSession(id); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"error":   false,
		"message": "success delete session by admin",
	})
}

func (s *SessionsController) GetSessionActive(ctx echo.Context) error {
	data, err := s.SessionService.GetSessionActive()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error":   true,
			"message": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"error":   false,
		"message": "success get all sessions",
		"data":    data,
	})
}
