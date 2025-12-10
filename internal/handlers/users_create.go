package handlers

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/yoshiyoshifujii/go-echo-sample/internal/api"
)

// CreateUser handles POST /users.
func (h *Handler) CreateUser(_ context.Context, request api.CreateUserRequestObject) (api.CreateUserResponseObject, error) {
	if request.Body == nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}

	req := request.Body
	if req.Name == "" || req.Email == "" {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "name and email are required")
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

	return api.CreateUser201JSONResponse(user), nil
}
