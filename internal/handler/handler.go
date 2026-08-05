package handler

import (
	"log/slog"

	"github.com/lmu3rto/exchange-platform/internal/service/contracts"
)

type Handler struct {
	userService contracts.UserRepository
	logger      *slog.Logger
}

func NewHandler(us contracts.UserRepository, logger *slog.Logger) *Handler {
	return &Handler{
		userService: us,
		logger:      logger,
	}
}
