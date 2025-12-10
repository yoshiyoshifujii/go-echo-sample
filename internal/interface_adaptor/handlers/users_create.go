package handlers

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	openapi_types "github.com/oapi-codegen/runtime/types"

	"github.com/yoshiyoshifujii/go-echo-sample/internal/interface_adaptor/api"
)

// CreateUser handles POST /users.
func (h *Handler) CreateUser(ctx context.Context, request api.CreateUserRequestObject) (api.CreateUserResponseObject, error) {
	if request.Body == nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}

	req := request.Body
	if req.Name == "" || req.Email == "" {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "name and email are required")
	}

	user, err := h.createUser.Execute(ctx, req.Name, string(req.Email))
	if err != nil {
		return nil, err
	}

	return api.CreateUser201JSONResponse(api.User{
		Id:    user.ID,
		Name:  user.Name,
		Email: openapi_types.Email(user.Email),
	}), nil
}
