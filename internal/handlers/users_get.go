package handlers

import (
	"context"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	openapi_types "github.com/oapi-codegen/runtime/types"

	"github.com/yoshiyoshifujii/go-echo-sample/internal/api"
	"github.com/yoshiyoshifujii/go-echo-sample/internal/repository"
)

// GetUser handles GET /users/{id}.
func (h *Handler) GetUser(ctx context.Context, request api.GetUserRequestObject) (api.GetUserResponseObject, error) {
	user, err := h.findUser.Execute(ctx, request.Id)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return nil, echo.NewHTTPError(http.StatusNotFound, "user not found")
		}
		return nil, err
	}

	return api.GetUser200JSONResponse(api.User{
		Id:    user.ID,
		Name:  user.Name,
		Email: openapi_types.Email(user.Email),
	}), nil
}
