// Package handlers contains Echo handlers implementing the generated ServerInterface.
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

// NewHandler returns a handler with in-memory storage.
func NewHandler() *Handler {
	return &Handler{
		nextID: 1,
	}
}
