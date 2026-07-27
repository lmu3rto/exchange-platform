package handler

import (
	"github.com/lmu3rto/exchange-platform/internal/service"

)
type Handler struct {
	userService *service.UserService
}

func NewHandler(us *service.UserService) (*Handler) {
	return &Handler {
		userService: us,
	}
}
