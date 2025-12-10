package handlers

import (
	"context"

	"github.com/yoshiyoshifujii/go-echo-sample/internal/interface_adaptor/api"
)

// GetHealth handles the health check endpoint.
func (h *Handler) GetHealth(_ context.Context, _ api.GetHealthRequestObject) (api.GetHealthResponseObject, error) {
	return api.GetHealth200JSONResponse(api.Health{Status: "ok"}), nil
}
