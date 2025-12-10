package handlers

import (
	"context"

	"github.com/yoshiyoshifujii/go-echo-sample/internal/api"
)

// ListUsers handles GET /users.
func (h *Handler) ListUsers(_ context.Context, _ api.ListUsersRequestObject) (api.ListUsersResponseObject, error) {
	h.mu.Lock()
	defer h.mu.Unlock()

	users := make([]api.User, len(h.users))
	copy(users, h.users)

	return api.ListUsers200JSONResponse(users), nil
}
