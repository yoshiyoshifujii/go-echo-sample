package handlers

import (
	"net/http"
	"sync"

	"github.com/labstack/echo/v4"

	"github.com/yoshiyoshifujii/go-echo-sample/internal/api"
)

// Handler implements the generated ServerInterface.
type Handler struct {
	mu     sync.Mutex
	users  []api.User
	nextID int64
}

func NewHandler() *Handler {
	return &Handler{
		nextID: 1,
	}
}

func (h *Handler) GetHealth(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, api.Health{Status: "ok"})
}

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

func (h *Handler) ListUsers(ctx echo.Context) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	users := make([]api.User, len(h.users))
	copy(users, h.users)

	return ctx.JSON(http.StatusOK, users)
}
