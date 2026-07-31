package handler

import (
	"github.com/lmu3rto/exchange-platform/internal/service"
)

type Handler struct {
	userService service.UserRepository
}

func NewHandler(us service.UserRepository) *Handler {
	return &Handler{
		userService: us,
	}
}
