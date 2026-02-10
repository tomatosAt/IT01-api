package handler

import (
	"github.com/tomatosAt/IT01-api/module/front-end/ports"
	service "github.com/tomatosAt/IT01-api/module/front-end/service"
)

/**
API endpoint input/output controlling and validation
*/

type Handler struct {
	svc ports.Service
}

func NewHandler(svc *service.Service) *Handler {
	return &Handler{
		svc: svc,
	}
}
