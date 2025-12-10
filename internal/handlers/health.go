package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/yoshiyoshifujii/go-echo-sample/internal/api"
)

// GetHealth handles the health check endpoint.
func (h *Handler) GetHealth(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, api.Health{Status: "ok"})
}
