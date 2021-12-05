package grpcHandler

import "github.com/1makarov/go-dater/server/internal/services"

type Handler struct {
	services *services.Services
}

func NewHandler(services *services.Services) *Handler {
	return &Handler{
		services: services,
	}
}
