package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/yoshiyoshifujii/go-echo-sample/internal/api"
)

// CreateUser handles POST /users.
func (h *Handler) CreateUser(ctx echo.Context) error {
	var req api.CreateUserRequest
	if err := ctx.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}
	if req.Name == "" || req.Email == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "name and email are required")
	}

	h.mu.Lock()
	defer h.mu.Unlock()

	user := api.User{
		Id:    h.nextID,
		Name:  req.Name,
		Email: req.Email,
	}
	h.nextID++

	h.users = append(h.users, user)

	return ctx.JSON(http.StatusCreated, user)
}
