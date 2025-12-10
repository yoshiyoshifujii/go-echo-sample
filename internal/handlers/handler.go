package handlers

import (
	"sync"

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
