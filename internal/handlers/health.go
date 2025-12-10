package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/yoshiyoshifujii/go-echo-sample/internal/api"
)

// HealthHandler implements the generated server interface for health checks.
type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) GetHealth(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, api.Health{Status: "ok"})
}
