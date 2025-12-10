// Package handlers contains Echo handlers implementing the generated ServerInterface.
package handlers

import "github.com/yoshiyoshifujii/go-echo-sample/internal/usecase"

// Handler implements the generated ServerInterface.
type Handler struct {
	createUser *usecase.CreateUser
	findUser   *usecase.FindUser
}

// NewHandler wires user use cases.
func NewHandler(createUser *usecase.CreateUser, findUser *usecase.FindUser) *Handler {
	return &Handler{
		createUser: createUser,
		findUser:   findUser,
	}
}
