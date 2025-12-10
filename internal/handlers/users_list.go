package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/yoshiyoshifujii/go-echo-sample/internal/api"
)

// ListUsers handles GET /users.
func (h *Handler) ListUsers(ctx echo.Context) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	users := make([]api.User, len(h.users))
	copy(users, h.users)

	return ctx.JSON(http.StatusOK, users)
}
